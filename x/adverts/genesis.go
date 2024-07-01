package adverts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"remap/x/adverts/keeper"
	"remap/x/adverts/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the advert
	for _, elem := range genState.AdvertList {
		k.SetAdvert(ctx, elem)
	}

	// Set advert count
	k.SetAdvertCount(ctx, genState.AdvertCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AdvertList = k.GetAllAdvert(ctx)
	genesis.AdvertCount = k.GetAdvertCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
