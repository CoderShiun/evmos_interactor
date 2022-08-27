package main

import (
	"errors"
	"evmosInteractor/contracts/account"
	"evmosInteractor/contracts/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	s    contract.Sample
	e    contract.ERC20
	user *account.User
)

// contractList get the account list from evmosInteractor.
func contractList() {
	byteList, err := exec.Command("evmosInteractor", "contract", "list").Output()
	if err != nil {
		log.Fatal(err)
	}

	if len(byteList) != 37 {
		log.Fatal("get contract list error")
	}

	fmt.Println(string(byteList))
}

// deploySample deploys the sample contract on evmos chain by using evmosInteractor command.
func deploySample() {
	result, err := exec.Command("evmosInteractor", "contract", "deploy", "sample").Output()
	if err != nil {
		log.Fatal(err)
	}

	if _, err = os.Stat("./sample.address"); errors.Is(err, os.ErrNotExist) {
		log.Fatal("can not find sample deployed file.")
	}
	if err != nil {
		log.Fatal("deploy sample contract error.")
	}

	fmt.Println(string(result))
}

// getSampleVersion gets the sample smart contract version.
func getSampleVersion() {
	time.Sleep(2 * time.Second)
	user = account.NewUser()

	sampleData, err := os.ReadFile("./sample.address")
	if err != nil || len(sampleData) == 0 {
		log.Fatal("got no sample.address data, ", err)
		return
	}

	s = contract.Sample{
		User:            *user,
		ContractAddress: common.HexToAddress(string(sampleData)),
	}
	s.ContractInstance = s.GetContractInstance(s.ContractAddress)

	s.GetVersion()
}

// deploySample deploys the erc20 smart contract on evmos chain by using evmosInteractor command.
func deployErc20() {
	result, err := exec.Command("evmosInteractor", "contract", "deploy", "erc20").Output()
	if err != nil {
		log.Fatal(err)
	}

	if _, err = os.Stat("./erc20.address"); errors.Is(err, os.ErrNotExist) {
		log.Fatal("can not find erc20 deployed file.")
	}
	if err != nil {
		log.Fatal("deploy erc20 contract error.")
	}

	fmt.Println(string(result))
}

// erc20Mint mints 1000 erc20 tokens.
func erc20Mint() {
	time.Sleep(2 * time.Second)

	data, err := os.ReadFile("./erc20.address")
	if err != nil || len(data) == 0 {
		log.Fatal("got no sample.address data, ", err)
		return
	}

	e = contract.ERC20{
		User:            *user,
		ContractAddress: common.HexToAddress(string(data)),
	}
	e.ContractInstance = e.GetContractInstance(s.ContractAddress)

	fmt.Println("mint 1000 tokens")
	e.Mint("1000")
}

// erc20Burn destroys 500 erc20 tokens.
func erc20Burn() {
	fmt.Println("burn 500 tokens")
	e.Burn("500")
}

// erc20Tx sends 100 erc20 tokens to address 0x808658fcEd5f1a1Bfdd8AB26F8609566b101A6E6.
func erc20Tx() {
	fmt.Println("send 100 token to address 0x808658fcEd5f1a1Bfdd8AB26F8609566b101A6E6")
	e.Transfer("0x808658fcEd5f1a1Bfdd8AB26F8609566b101A6E6", "100")
}
