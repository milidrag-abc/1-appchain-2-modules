package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"remap/x/assets/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				AssetList: []types.Asset{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				AssetCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated asset",
			genState: &types.GenesisState{
				AssetList: []types.Asset{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid asset count",
			genState: &types.GenesisState{
				AssetList: []types.Asset{
					{
						Id: 1,
					},
				},
				AssetCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
