package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"remap/x/assets/types"
)

func CmdCreateAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-asset [asset-id] [version] [state] [vc-available] [asset-type] [owner] [address] [asset-size] [properties]",
		Short: "Create a new asset",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAssetId := args[0]
			argVersion, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argState := args[2]
			argVcAvailable, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}
			argAssetType := args[4]
			argOwner := args[5]
			argAddress := args[6]
			argAssetSize, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}
			argProperties := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAsset(clientCtx.GetFromAddress().String(), argAssetId, argVersion, argState, argVcAvailable, argAssetType, argOwner, argAddress, argAssetSize, argProperties)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-asset [id] [asset-id] [version] [state] [vc-available] [asset-type] [owner] [address] [asset-size] [properties]",
		Short: "Update a asset",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argAssetId := args[1]

			argVersion, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			argState := args[3]

			argVcAvailable, err := cast.ToBoolE(args[4])
			if err != nil {
				return err
			}

			argAssetType := args[5]

			argOwner := args[6]

			argAddress := args[7]

			argAssetSize, err := cast.ToUint64E(args[8])
			if err != nil {
				return err
			}

			argProperties := args[9]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAsset(clientCtx.GetFromAddress().String(), id, argAssetId, argVersion, argState, argVcAvailable, argAssetType, argOwner, argAddress, argAssetSize, argProperties)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-asset [id]",
		Short: "Delete a asset by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAsset(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
