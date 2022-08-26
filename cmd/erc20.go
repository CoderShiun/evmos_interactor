package cmd

import (
	"bufio"
	"evmosInteractor/contracts/account"
	"evmosInteractor/contracts/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const (
	TotalSupply = "total_supply"
	Decimals    = "decimals"
	GetBalance  = "get_balance"
)

var erc20 contract.ERC20

// erc20Cmd represents the erc20 command
var erc20Cmd = &cobra.Command{
	Use:   "erc20",
	Short: "erc20 contract functions",
	Long:  `it provides additional functions for the erc20 smart contract.`,
	Run: func(cmd *cobra.Command, args []string) {
		readErc20ContractInfo()

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("total_supply, decimals, get_balance")
		fmt.Println("which function are you going to use?")
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		f := strings.Join(strings.Fields(text), "")
		text = strings.ToLower(f)

		switch text {
		case TotalSupply:
			erc20.GetTotalSupply()
		case Decimals:
			erc20.GetDecimals()
		case GetBalance:
			reader = bufio.NewReader(os.Stdin)
			fmt.Println("please enter the address")
			fmt.Print("-> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			addr := strings.Join(strings.Fields(input), "")

			erc20.BalanceOf(common.HexToAddress(addr))

		default:
			fmt.Println("function is not existed.")
		}
	},
}

func readErc20ContractInfo() {
	data, err := os.ReadFile("./erc20.address")
	if err != nil || len(data) == 0 {
		log.Fatal("please deploy erc20 contract first, ", err)
		return
	}

	user := account.NewUser()
	erc20 = contract.ERC20{User: *user}

	erc20.ContractAddress = common.HexToAddress(string(data))
	erc20.ContractInstance = erc20.GetContractInstance(erc20.ContractAddress)
}
