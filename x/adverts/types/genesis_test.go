package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"remap/x/adverts/types"
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

				AdvertList: []types.Advert{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				AdvertCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated advert",
			genState: &types.GenesisState{
				AdvertList: []types.Advert{
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
			desc: "invalid advert count",
			genState: &types.GenesisState{
				AdvertList: []types.Advert{
					{
						Id: 1,
					},
				},
				AdvertCount: 0,
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
