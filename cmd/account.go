package cmd

import (
	"evmosInteractor/contracts/account"
	"fmt"
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "commands for your evmos account",
	Long:  `here you are able to check the account status`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "show your account balance",
	Long:  `it shows the balance from the evmos local test network account`,
	Run: func(cmd *cobra.Command, args []string) {
		user := account.NewUser()
		balance := user.GetBalance()
		fmt.Println(fmt.Sprintf("Balance: %v evm", balance))
	},
}
