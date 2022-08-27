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

// DeployContract deploys the erc20 smart contract.
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

// GetContractInstance returns the erc20 smart contract instance.
func (e *ERC20) GetContractInstance(contractAddress common.Address) *erc20.Contracts {
	instance, err := erc20.NewContracts(contractAddress, e.User.Cli)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

// GetTotalSupply returns the total supply of the erc20 token.
func (e *ERC20) GetTotalSupply() {
	total, err := e.ContractInstance.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total supply: ", total)
}

// GetDecimals returns the decimals of the erc20 smart contract.
func (e *ERC20) GetDecimals() {
	deci, err := e.ContractInstance.Decimals(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("decimals: ", deci)
}

// BalanceOf returns the balance of the erc20 smart contract account.
func (e *ERC20) BalanceOf(addr common.Address) {
	balance, err := e.ContractInstance.BalanceOf(nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("erc20 tokens balance: ", balance)
}

// Mint creates amount of new erc20 tokens, it increases the total supply.
func (e *ERC20) Mint(amount string) {
	tx, err := e.ContractInstance.Mint(e.User.GetAuth(), utils.GetBigInt(amount))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mint token successful, tx: ", tx.Hash())
}

// Burn burns amount of erc20 tokens, it decreases the total supply.
func (e *ERC20) Burn(amount string) {
	tx, err := e.ContractInstance.Burn(e.User.GetAuth(), utils.GetBigInt(amount))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("burn token successful, tx: ", tx.Hash())
}

// Transfer transfers amount of erc20 tokens to another account.
func (e *ERC20) Transfer(addr string, amount string) {
	tx, err := e.ContractInstance.Transfer(e.User.GetAuth(), common.HexToAddress(addr), utils.GetBigInt(amount))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("send tx successful, tx: ", tx.Hash())
}
