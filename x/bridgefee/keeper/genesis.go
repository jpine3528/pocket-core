package keeper

import (
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis - Init store state from genesis data
func (k Keeper) InitGenesis(ctx sdk.Ctx, data types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)

	for _, tokenInfo := range data.TokenInfos {
		k.SetTokenInfo(ctx, tokenInfo)
	}
	for _, target := range data.TokenTargets {
		k.SetTokenTargetInfo(ctx, target)
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns a GenesisState for a given context and keeper
func (k Keeper) ExportGenesis(ctx sdk.Ctx) types.GenesisState {
	return types.NewGenesisState(
		k.GetParams(ctx),
		k.GetAllTokenInfos(ctx),
		k.GetAllTokenTargetInfos(ctx),
	)
}
