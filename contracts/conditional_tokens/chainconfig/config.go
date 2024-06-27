package chainconfig

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type ChainConfig struct {
	Address common.Address
}

func GetChainConfig(chainId int) (*ChainConfig, error) {
	switch chainId {
	case 137:
		return &ChainConfig{
			Address: common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
		}, nil
	case 80002:
		return &ChainConfig{
			Address: common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
		}, nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Invalid chainId: %d", chainId))
	}
}
