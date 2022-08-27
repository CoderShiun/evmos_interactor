package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// contractCmd represents the contract command,
// it is used to deploy and interact with the smart contracts.
var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "contract progress",
	Long:  `interact with smart contracts`,
}

// listCmd shows the list of the smart contract.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show contract list",
	Long:  `show all the current smart contracts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("contract 1: Sample")
		fmt.Println("contract 2: ERC20")
	},
}
