package cmd

import (
	"context"
	"evmosInteractor/contracts/account"
	"evmosInteractor/logger"

	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "evmosInteractor",
	Short: "Testing on evmos chain",
	Long: `Deploy smart contracts and sand transactions on evmos chain, 
 it based on geth library.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		/*if v, err := cmd.Flags().GetBool("debug"); v || err != nil {
			fmt.Println("Setting debug level")
			logger.SetLevel(logger.DebugLevel)
		}*/
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// init initializes public variables and sets commands.
func init() {
	account.User = account.NewUser()
	initIPFS()

	rootCmd.AddCommand(contractCmd)
	rootCmd.AddCommand(accountCmd)

	accountCmd.AddCommand(balanceCmd)
	accountCmd.AddCommand(balanceOfCmd)
	accountCmd.AddCommand(transferCmd)

	contractCmd.AddCommand(listCmd)
	contractCmd.AddCommand(deployCmd)
	contractCmd.AddCommand(erc20Cmd)
	contractCmd.AddCommand(sampleCmd)
}

func initIPFS() {
	ctx := context.Background()

	state, _ := logger.Sh.FilesStat(ctx, "/evmosInter/eth")
	if state == nil {
		err := logger.Sh.FilesMkdir(ctx, "/evmosInter/eth", func(builder *shell.RequestBuilder) error {
			builder.Option("parents", true)
			return nil
		})
		if err != nil {
			log.Error(err)
		}

		err = logger.Sh.FilesMkdir(ctx, "/evmosInter/erc20")
		if err != nil {
			log.Error(err)
		}

		err = logger.Sh.FilesMkdir(ctx, "/evmosInter/sample")
		if err != nil {
			log.Error(err)
		}
	}
}
