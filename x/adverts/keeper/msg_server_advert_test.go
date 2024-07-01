package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"remap/x/adverts/types"
)

func TestAdvertMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAdvert(ctx, &types.MsgCreateAdvert{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestAdvertMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateAdvert
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAdvert{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAdvert{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAdvert{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateAdvert(ctx, &types.MsgCreateAdvert{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAdvert(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAdvertMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteAdvert
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAdvert{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAdvert{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAdvert{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateAdvert(ctx, &types.MsgCreateAdvert{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAdvert(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
