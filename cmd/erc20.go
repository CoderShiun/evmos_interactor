package cmd

import (
	"evmosInteractor/contracts/account"
	"evmosInteractor/contracts/contract"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	TotalSupply = "total_supply"
	Decimals    = "decimals"
	GetBalance  = "get_balance"
	Mint        = "mint"
	Burn        = "burn"
	Transfer    = "transfer"
)

var erc20 contract.ERC20

// erc20Cmd represents the erc20 command,
// it provides additional functions for the erc20 smart contract.
var erc20Cmd = &cobra.Command{
	Use:   "erc20",
	Short: "erc20 contract functions",
	Long:  `it provides additional functions for the erc20 smart contract.`,
	Run: func(cmd *cobra.Command, args []string) {
		readErc20ContractInfo()

		switch ercBasic() {
		case TotalSupply:
			fmt.Println("total supply: ", erc20.GetTotalSupply())
		case Decimals:
			fmt.Println("decimals: ", erc20.GetDecimals())
		case GetBalance:
			fmt.Println("erc20 tokens balance: ", erc20.BalanceOf(common.HexToAddress(ercGetBalance())))
		case Mint:
			erc20.Mint(ercMint())
		case Burn:
			erc20.Burn(ercBurn())
		case Transfer:
			erc20.Transfer(ercTransfer())
		default:
			fmt.Println("function is not existed.")
		}
	},
}

// readErc20ContractInfo reads the erc20 contract address from the file,
// and initializes the erc20 contract struct.
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
