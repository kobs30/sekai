package cli

import (
	"github.com/KiraCore/sekai/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strconv"
)

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "bridge sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(TxChangeCosmosEtherium())

	return txCmd
}

func TxChangeCosmosEtherium() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change",
		Short: "Create new change request from Cosmos to Etherium",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			to := args[1]

			inAmount, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			outAmount, _ := strconv.ParseInt(args[3], 10, 64)

			msg := types.NewMsgChangeCosmosEtherium(
				clientCtx.FromAddress,
				to,
				inAmount,
				outAmount,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(OldKey, "", "Previous hash string.")
	cmd.MarkFlagRequired(OldKey)

	cmd.Flags().String(NewKey, "", "Next hash string.")
	cmd.MarkFlagRequired(NewKey)

	cmd.Flags().String(NextAddress, "", "Next address to control the settings.")
	cmd.Flags().String(TargetAddress, "", "Target of the control request.")

	flags.AddTxFlagsToCmd(cmd)
	cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
