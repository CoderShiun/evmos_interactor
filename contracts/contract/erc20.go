package contract

import (
	"evmosInteractor/contracts/account"
	"evmosInteractor/contracts/erc20"
	"evmosInteractor/utils"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type ERC20 struct {
	User             account.User
	ContractInstance *erc20.Contracts
	ContractAddress  common.Address
}

func (e *ERC20) DeployContract() {
	address, tx, instance, err := erc20.DeployContracts(e.User.GetAuth(), e.User.Cli)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successful, deploy contract tx: ", tx.Hash().Hex())

	e.ContractInstance = instance
	e.ContractAddress = address

	f := []byte(address.String())
	err = os.WriteFile("./erc20.address", f, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (e *ERC20) GetContractInstance(contractAddress common.Address) *erc20.Contracts {
	instance, err := erc20.NewContracts(contractAddress, e.User.Cli)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

func (e *ERC20) GetTotalSupply() {
	total, err := e.ContractInstance.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total supply: ", total)
}

func (e *ERC20) GetDecimals() {
	deci, err := e.ContractInstance.Decimals(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("decimals: ", deci)
}

func (e *ERC20) BalanceOf(addr common.Address) {
	balance, err := e.ContractInstance.BalanceOf(nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("erc20 tokens balance: ", balance)
}

func (e *ERC20) Mint(amount string) {
	tx, err := e.ContractInstance.Mint(e.User.GetAuth(), utils.GetBigInt(amount))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mint token successful, tx: ", tx.Hash())
}

func (e *ERC20) Burn(amount string) {
	tx, err := e.ContractInstance.Burn(e.User.GetAuth(), utils.GetBigInt(amount))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("burn token successful, tx: ", tx.Hash())
}

func (e *ERC20) Transfer(addr string, amount string) {
	tx, err := e.ContractInstance.Transfer(e.User.GetAuth(), common.HexToAddress(addr), utils.GetBigInt(amount))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("send tx successful, tx: ", tx.Hash())
}
