package chainconfig

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type ChainConfig struct {
	OptimisticOracleV1 common.Address
	OptimisticOracleV2 common.Address
}

func GetChainConfig(chainId int) (*ChainConfig, error) {
	switch chainId {
	case 137:
		return &ChainConfig{
			OptimisticOracleV1: common.HexToAddress("0xBb1A8db2D4350976a11cdfA60A1d43f97710Da49"),
			OptimisticOracleV2: common.HexToAddress("0xee3afe347d5c74317041e2618c49534daf887c24"),
		}, nil
	case 80002:
		return &ChainConfig{
			OptimisticOracleV1: common.HexToAddress(""),
			OptimisticOracleV2: common.HexToAddress("0x38fAc33bD20D4c4Cce085C0f347153C06CbA2968"),
		}, nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("Invalid chainId: %d", chainId))
	}
}
