package cmd

import (
	"github.com/spf13/cobra"
)

// contractCmd represents the contract command
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "contract progress",
	Long:  `interact with smart contracts`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(contractCmd)
}
