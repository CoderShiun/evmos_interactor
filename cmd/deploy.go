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

// deployCmd deploys the smart contracts.
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy a new contract",
	Long:  `choose one of the contract to deploy on the evmos chain`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please give a contract name e.g. erc20")
		}

		for _, v := range args {
			v = strings.ToLower(v)
			switch v {
			case ERC20:
				erc20 := contract.ERC20{
					User: *account.User,
				}
				contract.DeployContract(&erc20)
				//erc20.Deploy()

			case Sample:
				sample := contract.Sample{
					User: *account.User,
				}
				contract.DeployContract(&sample)
				//sample.Deploy()

			default:
				fmt.Println(fmt.Sprintf("%v is not a valid contract name", v))
			}
		}
	},
}
