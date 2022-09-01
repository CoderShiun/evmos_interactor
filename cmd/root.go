package cmd

import (
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
			log.SetLevel(log.DebugLevel)
		}*/
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// init initializes the config and sets commands.
func init() {
	rootCmd.AddCommand(contractCmd)
	rootCmd.AddCommand(accountCmd)

	accountCmd.AddCommand(balanceCmd)

	contractCmd.AddCommand(listCmd)
	contractCmd.AddCommand(deployCmd)
	contractCmd.AddCommand(erc20Cmd)
	contractCmd.AddCommand(sampleCmd)
}
