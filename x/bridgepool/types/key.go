package types

import (
	sdk "github.com/pokt-network/pocket-core/types"
)

const (
	ModuleName   = "bridgepool" // name of the module
	RouterKey    = ModuleName   // RouterKey defines the routing key for a Parameter Change
	StoreKey     = ModuleName   // key for state store
	TStoreKey    = "transient_bridgepool"
	QuerierRoute = ModuleName // QuerierRoute is the querier route for the bridgepool module
)

var (
	FeeRateKeyPrefix       = []byte{0x00}
	AllowedTargetKeyPrefix = []byte{0x01}
	LiquidityKeyPrefix     = []byte{0x02}
	SignerKeyPrefix        = []byte{0x03}
	WithdrawSaltKeyPrefix  = []byte{0x04}
)

func FeeRateKey(token string) []byte {
	return append(FeeRateKeyPrefix, token...)
}

func AllowedTargetKey(token string, chainId string) []byte {
	return append(append(AllowedTargetKeyPrefix, token...), []byte(chainId)...)
}

func LiquidityKey(token string, user sdk.Address) []byte {
	return append(append(LiquidityKeyPrefix, token...), user...)
}

func SignerKey(signer string) []byte {
	return append(SignerKeyPrefix, signer...)
}

func WithdrawSaltKey(salt []byte) []byte {
	return append(WithdrawSaltKeyPrefix, salt...)
}
