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
	Version = "version"
	SetItem = "set_item"
	GetItem = "get_item"
)

var s contract.Sample

// sampleCmd represents the sample command
var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "sample contract functions",
	Long:  `it provides additional functions for the sample smart contract.`,
	Run: func(cmd *cobra.Command, args []string) {
		readSampleContractInfo()

		switch sampBasic() {
		case Version:
			s.GetVersion()
		case SetItem:
			s.SetItem(sampSetItem())
		case GetItem:
			s.GetItems(sampGetItem())
		default:
			fmt.Println("function is not existed.")
		}
	},
}

func readSampleContractInfo() {
	data, err := os.ReadFile("./sample.address")
	if err != nil || len(data) == 0 {
		log.Fatal("please deploy sample contract first, ", err)
	}

	user := account.NewUser()
	s = contract.Sample{User: *user}
	s.ContractAddress = common.HexToAddress(string(data))
	s.ContractInstance = s.GetContractInstance(s.ContractAddress)
}
