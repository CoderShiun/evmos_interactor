package cmd

import (
	"bufio"
	"evmosInteractor/contracts/account"
	"evmosInteractor/contracts/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

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

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("version, set_item, get_item")
		fmt.Println("which function are you going to use?")
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		f := strings.Join(strings.Fields(text), "")
		text = strings.ToLower(f)

		switch text {
		case Version:
			s.GetVersion()
		case SetItem:
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("please give key")
			fmt.Print("-> ")
			key, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			reader = bufio.NewReader(os.Stdin)
			fmt.Println("please give value")
			fmt.Print("-> ")
			value, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			s.SetItem(key, value)
		case GetItem:
			reader = bufio.NewReader(os.Stdin)
			fmt.Println("which item are you searching for?")
			fmt.Print("-> ")
			key, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			s.GetItems(key)
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
