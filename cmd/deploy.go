package cmd

import (
	"evmosInteractor/contracts/account"
	"evmosInteractor/contracts/contract"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	Sample = "sample"
	ERC20  = "erc20"
)

// deployCmd represents the related deployment command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a new contract",
	Long:  `choose one of the contract to deploy on the evmos chain`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please give a contract name e.g. erc20")
		}

		user := account.NewUser()

		for _, v := range args {
			v = strings.ToLower(v)
			switch v {
			case ERC20:
				erc20 := contract.ERC20{
					User: *user,
				}
				erc20.DeployContract()

			case Sample:
				sample := contract.Sample{
					User: *user,
				}
				sample.DeployContract()

			default:
				fmt.Println(fmt.Sprintf("%v is not a correct name"))
				return
			}
		}
	},
}

func init() {
	contractCmd.AddCommand(deployCmd)
}