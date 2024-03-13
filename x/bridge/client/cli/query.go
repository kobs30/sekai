package cli

import (
	"context"
	"github.com/KiraCore/sekai/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewQueryCmd returns a root CLI command handler for all x/distributor transaction commands.
func NewQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:   types.RouterKey,
		Short: "query commands for the bridge module",
	}

	queryCmd.AddCommand(GetCmdQueryChangeCosmosEtheriumByAddress())
	queryCmd.AddCommand(GetCmdQueryChangeEtheriumCosmosByAddress())

	return queryCmd
}

// GetCmdQueryChangeCosmosEtheriumByAddress is the querier for change by address.
func GetCmdQueryChangeCosmosEtheriumByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get_cosmos_etherium [addr]",
		Short: "Query change from Cosmos to Etherium assigned to an address",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			accAddr, err := sdk.AccAddressFromBech32(args[0])

			if err != nil {
				return errors.Wrap(err, "invalid account address")
			}

			params := &types.ChangeCosmosEtheriumByAddressRequest{Addr: accAddr}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ChangeCosmosEtheriumByAddress(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryChangeEtheriumCosmosByAddress is the querier for change by address.
func GetCmdQueryChangeEtheriumCosmosByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get_etherium_cosmos [addr]",
		Short: "Query change from Etherium to Cosmos assigned to an address",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			accAddr, err := sdk.AccAddressFromBech32(args[0])

			if err != nil {
				return errors.Wrap(err, "invalid account address")
			}

			params := &types.ChangeEtheriumCosmosByAddressRequest{Addr: accAddr}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ChangeEtheriumCosmosByAddress(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
