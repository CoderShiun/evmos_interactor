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

// balanceCmd returns the evmos token balance of the account
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "show your account balance",
	Long:  `it shows the balance from the evmos local test network account`,
	Run: func(cmd *cobra.Command, args []string) {
		balance := account.User.GetBalance()
		fmt.Println(fmt.Sprintf("Balance: %v evm", balance))
	},
}

// balanceOfCmd returns the evmos balance from a specific account
var balanceOfCmd = &cobra.Command{
	Use:   "balanceOf",
	Short: "get balance from an account",
	Long:  `get balance from a given account`,
	Run: func(cmd *cobra.Command, args []string) {
		balance := account.User.GetBalanceOf(ercGetBalance())
		fmt.Println(fmt.Sprintf("Balance: %v evm", balance))
	},
}

// transferCmd returns the evmos token balance of the account
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "send evmos token",
	Long:  `send evmos to a specific account`,
	Run: func(cmd *cobra.Command, args []string) {
		account.User.SendEVMOS(ercTransfer())
	},
}
