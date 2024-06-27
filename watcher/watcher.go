package watcher

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/playground/go-market-events/constants"
	"github.com/playground/go-market-events/contracts/conditional_tokens"
	optimistic_oracle_chainconfig "github.com/playground/go-market-events/contracts/optimistic_oracle/chainconfig"
	optimistic_oracle_v1 "github.com/playground/go-market-events/contracts/optimistic_oracle/v1"
	optimistic_oracle_v2 "github.com/playground/go-market-events/contracts/optimistic_oracle/v2"
	"github.com/sirupsen/logrus"
)

func ListenEvents(ctx context.Context, client *ethclient.Client, filter ethereum.FilterQuery) error {
	logs := make(chan types.Log)
	subs, err := client.SubscribeFilterLogs(ctx, filter, logs)
	if err != nil {
		return fmt.Errorf("error subscribing to filter logs: %s", err.Error())
	}

	for {
		select {
		case err, ok := <-subs.Err():
			if ok && err != nil {
				close(logs)
				subs.Unsubscribe()
				return err
			}
			return nil
		case log, ok := <-logs:
			if !ok {
				close(logs)
				subs.Unsubscribe()
				return fmt.Errorf("logs channel closed")
			}

			if err := processLog(ctx, log); err != nil {
				logrus.Errorf("error parsing log: %s", err.Error())
			}
		}
	}
}

func processLog(_ context.Context, log types.Log) error {
	logName, ok := constants.TOPICS[log.Topics[0].Hex()]
	if !ok {
		return fmt.Errorf("error processing event - invalid event")
	}

	chainCfg, err := optimistic_oracle_chainconfig.GetChainConfig(int(constants.CHAIN_ID))
	if err != nil {
		return err
	}

	switch logName {
	case constants.PROPOSE_PRICE:
		if log.Address.Cmp(chainCfg.OptimisticOracleV1) == 0 {
			parsed, err := getOptimisticOracleV1Contract(log.Address).ParseProposePrice(log)
			if err != nil {
				return err
			}

			logrus.Infof("new %s V1 event, question id: %s", logName, crypto.Keccak256Hash(parsed.AncillaryData))
		} else {
			parsed, err := getOptimisticOracleV2Contract(log.Address).ParseProposePrice(log)
			if err != nil {
				return err
			}

			logrus.Infof("new %s V2 event, question id: %s", logName, crypto.Keccak256Hash(parsed.AncillaryData))
		}

	case constants.DISPUTE_PRICE:
		if log.Address.Cmp(chainCfg.OptimisticOracleV1) == 0 {
			parsed, err := getOptimisticOracleV1Contract(log.Address).ParseDisputePrice(log)
			if err != nil {
				return err
			}

			logrus.Infof("new %s V1 event, question id: %s", logName, crypto.Keccak256Hash(parsed.AncillaryData))
		} else {
			parsed, err := getOptimisticOracleV2Contract(log.Address).ParseDisputePrice(log)
			if err != nil {
				return err
			}

			logrus.Infof("new %s V2 event, question id: %s", logName, crypto.Keccak256Hash(parsed.AncillaryData))
		}

	case constants.CONDITIONAL_TOKENS_CONDITION_RESOLUTION:
		parsed, err := getConditionalTokensContract(log.Address).ParseConditionResolution(log)
		if err != nil {
			return err
		}

		logrus.Infof("new %s event, question id: %s", logName, parsed.QuestionId)

	default:
		return fmt.Errorf("unknown log %s", logName)
	}

	return nil
}

func getOptimisticOracleV1Contract(address common.Address) *optimistic_oracle_v1.OptimisticOracleV1Filterer {
	i, _ := optimistic_oracle_v1.NewOptimisticOracleV1Filterer(address, nil)
	return i
}

func getOptimisticOracleV2Contract(address common.Address) *optimistic_oracle_v2.OptimisticOracleV2Filterer {
	i, _ := optimistic_oracle_v2.NewOptimisticOracleV2Filterer(address, nil)
	return i
}

func getConditionalTokensContract(address common.Address) *conditional_tokens.ConditionalTokensFilterer {
	i, _ := conditional_tokens.NewConditionalTokensFilterer(address, nil)
	return i
}
