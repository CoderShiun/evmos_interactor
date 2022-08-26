package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show contract list",
	Long:  `show all the current smart contracts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("contract 1: Sample")
		fmt.Println("contract 2: ERC20")
	},
}

func init() {
	contractCmd.AddCommand(listCmd)
}
