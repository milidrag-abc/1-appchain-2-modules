package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "remap/testutil/keeper"
	"remap/x/assets/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AssetsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
