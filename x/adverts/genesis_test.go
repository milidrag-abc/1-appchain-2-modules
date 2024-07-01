package adverts_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "remap/testutil/keeper"
	"remap/testutil/nullify"
	"remap/x/adverts"
	"remap/x/adverts/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AdvertList: []types.Advert{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AdvertCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AdvertsKeeper(t)
	adverts.InitGenesis(ctx, *k, genesisState)
	got := adverts.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AdvertList, got.AdvertList)
	require.Equal(t, genesisState.AdvertCount, got.AdvertCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
