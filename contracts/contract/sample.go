package contract

import (
	"evmosInteractor/contracts/account"
	"evmosInteractor/contracts/sample"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type Sample struct {
	User             account.User
	ContractInstance *sample.Contracts
	ContractAddress  common.Address
}

// Deploy deploys the sample smart contract.
func (s *Sample) Deploy() {
	input := "1.0"
	address, tx, instance, err := sample.DeployContracts(s.User.GetAuth(), s.User.Cli, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successful, deploy contract tx: ", tx.Hash().Hex())

	s.ContractInstance = instance
	s.ContractAddress = address

	f := []byte(address.String())
	err = os.WriteFile("./sample.address", f, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// GetContractInstance returns the sample smart contract instance.
func (s *Sample) GetContractInstance(contractAddress common.Address) *sample.Contracts {
	instance, err := sample.NewContracts(contractAddress, s.User.Cli)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

// GetVersion returns the version of the sample smart contract.
func (s *Sample) GetVersion() {
	v, err := s.ContractInstance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("sample contract version: %v", v))
}

// SetItem sets the item of the sample smart contract.
func (s *Sample) SetItem(itemKey, itemValue string) {
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], itemKey)
	copy(value[:], itemValue)

	tx, err := s.ContractInstance.SetItem(s.User.GetAuth(), key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("set item tx: %s", tx.Hash().Hex()))
}

// GetItems returns the value from the sample smart contract item.
func (s *Sample) GetItems(name string) {
	key := [32]byte{}
	copy(key[:], name)

	result, err := s.ContractInstance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("item value: ", string(result[:]))
}
