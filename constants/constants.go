package constants

import "github.com/ethereum/go-ethereum/crypto"

const CHAIN_ID = 137

const (
	PROPOSE_PRICE = "ProposePrice(address,address,bytes32,uint256,bytes,int256,uint256,address)"
	DISPUTE_PRICE = "DisputePrice(address,address,address,bytes32,uint256,bytes,int256)"

	CONDITIONAL_TOKENS_CONDITION_RESOLUTION = "ConditionResolution(bytes32,address,bytes32,uint256,uint256[])"
)

var (
	PROPOSE_PRICE_SIG_HASH = crypto.Keccak256Hash([]byte(PROPOSE_PRICE))
	DISPUTE_PRICE_SIG_HASH = crypto.Keccak256Hash([]byte(DISPUTE_PRICE))

	CONDITIONAL_TOKENS_CONDITION_RESOLUTION_SIG_HASH = crypto.Keccak256Hash([]byte(CONDITIONAL_TOKENS_CONDITION_RESOLUTION))
)

// topics
var TOPICS = map[string]string{
	PROPOSE_PRICE_SIG_HASH.Hex():                           PROPOSE_PRICE,
	DISPUTE_PRICE_SIG_HASH.Hex():                           DISPUTE_PRICE,
	CONDITIONAL_TOKENS_CONDITION_RESOLUTION_SIG_HASH.Hex(): CONDITIONAL_TOKENS_CONDITION_RESOLUTION,
}
