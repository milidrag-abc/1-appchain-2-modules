package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"remap/x/adverts/types"
)

func CmdCreateAdvert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-advert [advert-id] [state] [asset-id] [price] [min-rental-period] [max-rental-period] [conditions] [expiration-date]",
		Short: "Create a new advert",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAdvertId := args[0]
			argState := args[1]
			argAssetId := args[2]
			argPrice, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argMinRentalPeriod, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argMaxRentalPeriod, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argConditions := args[6]
			argExpirationDate, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAdvert(clientCtx.GetFromAddress().String(), argAdvertId, argState, argAssetId, argPrice, argMinRentalPeriod, argMaxRentalPeriod, argConditions, argExpirationDate)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAdvert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-advert [id] [advert-id] [state] [asset-id] [price] [min-rental-period] [max-rental-period] [conditions] [expiration-date]",
		Short: "Update a advert",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argAdvertId := args[1]

			argState := args[2]

			argAssetId := args[3]

			argPrice, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			argMinRentalPeriod, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			argMaxRentalPeriod, err := cast.ToUint64E(args[6])
			if err != nil {
				return err
			}

			argConditions := args[7]

			argExpirationDate, err := cast.ToUint64E(args[8])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAdvert(clientCtx.GetFromAddress().String(), id, argAdvertId, argState, argAssetId, argPrice, argMinRentalPeriod, argMaxRentalPeriod, argConditions, argExpirationDate)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAdvert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-advert [id]",
		Short: "Delete a advert by id",
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

			msg := types.NewMsgDeleteAdvert(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
