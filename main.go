package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/playground/go-market-events/constants"
	conditional_tokens_chainconfig "github.com/playground/go-market-events/contracts/conditional_tokens/chainconfig"
	optimistic_oracle_chainconfig "github.com/playground/go-market-events/contracts/optimistic_oracle/chainconfig"
	"github.com/playground/go-market-events/logger"
	"github.com/playground/go-market-events/watcher"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.SetLogger("trace")
	ctx := context.Background()
	rpcUrl := "wss://polygon-bor-rpc.publicnode.com" // you can use a better one, this one is public
	client, err := ethclient.DialContext(ctx, rpcUrl)
	if err != nil {
		logrus.Fatalf("Error DialContext: %s", err.Error())
	}

	// when a market is ready to be resolved, someone is going to propose an outcome
	go listenProposePriceEvents(ctx, client)

	// if the resolution is not accepted, someone is going to dispute it
	go listenDisputePriceEvents(ctx, client)

	// when the final proposal is ready, the market resolves
	go listenConditionResolutionEvents(ctx, client)

	waitUntilStop()
}

func waitUntilStop() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals // wait for signal

	go func() {
		<-signals // hard exit on second signal (in case shutdown gets stuck)
		os.Exit(1)
	}()
}

func listenProposePriceEvents(ctx context.Context, client *ethclient.Client) {
	chainCfg, err := optimistic_oracle_chainconfig.GetChainConfig(int(constants.CHAIN_ID))
	if err != nil {
		logrus.Fatalf("Error getting chainconfig: %s", err.Error())
	}

	filter := ethereum.FilterQuery{
		Addresses: []common.Address{chainCfg.OptimisticOracleV1, chainCfg.OptimisticOracleV2},
		Topics: [][]common.Hash{
			{constants.PROPOSE_PRICE_SIG_HASH},
			{}, // Any from address allowed
		},
	}

	for {
		if err := watcher.ListenEvents(ctx, client, filter); err != nil {
			logrus.Warnf("error on ListenEvents: %s", err.Error())
		}
	}
}

func listenDisputePriceEvents(ctx context.Context, client *ethclient.Client) {
	chainCfg, err := optimistic_oracle_chainconfig.GetChainConfig(int(constants.CHAIN_ID))
	if err != nil {
		logrus.Fatalf("Error getting chainconfig: %s", err.Error())
	}

	filter := ethereum.FilterQuery{
		Addresses: []common.Address{chainCfg.OptimisticOracleV1, chainCfg.OptimisticOracleV2},
		Topics: [][]common.Hash{
			{constants.DISPUTE_PRICE_SIG_HASH},
			{}, // Any from address allowed
		},
	}

	for {
		if err := watcher.ListenEvents(ctx, client, filter); err != nil {
			logrus.Warnf("error on ListenEvents: %s", err.Error())
		}
	}
}

func listenConditionResolutionEvents(ctx context.Context, client *ethclient.Client) {
	chainCfg, err := conditional_tokens_chainconfig.GetChainConfig(int(constants.CHAIN_ID))
	if err != nil {
		logrus.Fatalf("Error getting chainconfig: %s", err.Error())
	}

	filter := ethereum.FilterQuery{
		Addresses: []common.Address{chainCfg.Address},
		Topics:    [][]common.Hash{{constants.CONDITIONAL_TOKENS_CONDITION_RESOLUTION_SIG_HASH}},
	}

	for {
		if err := watcher.ListenEvents(ctx, client, filter); err != nil {
			logrus.Warnf("error on ListenEvents: %s", err.Error())
		}
	}
}
