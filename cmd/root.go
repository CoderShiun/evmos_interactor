package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Interactor",
	Short: "Testing on evmos chain",
	Long: `Deploy smart contracts and sand transactions on evmos chain, 
 it based on geth library.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if v, err := cmd.Flags().GetBool("debug"); v || err != nil {
			fmt.Println("Setting debug level")
			log.SetLevel(log.DebugLevel)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//balance := user.GetBalance()
	//fmt.Println("balance: ", balance)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Evmos.yaml)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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
}
