package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "remap/testutil/keeper"
	"remap/x/adverts/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AdvertsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
