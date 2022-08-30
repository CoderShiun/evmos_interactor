package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
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
	//cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Evmos.yaml)")
	//rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug mode")

	rootCmd.AddCommand(contractCmd)
	rootCmd.AddCommand(accountCmd)

	accountCmd.AddCommand(balanceCmd)

	contractCmd.AddCommand(listCmd)
	contractCmd.AddCommand(deployCmd)
	contractCmd.AddCommand(erc20Cmd)
	contractCmd.AddCommand(sampleCmd)
}

// initConfig reads in config file and ENV variables if set.
/*func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".inter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".inter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}*/
