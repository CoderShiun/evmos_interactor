package contract

import (
	"evmosInteractor/contracts/account"
	"evmosInteractor/contracts/sample"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"os"
)

type Sample struct {
	User             account.User
	ContractInstance *sample.Contracts
	ContractAddress  common.Address
}

func (s *Sample) DeployContract() {
	input := "1.0"
	address, tx, instance, err := sample.DeployContracts(s.User.GetAuth(), s.User.Cli, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successful, deploy contract tx: ", tx.Hash().Hex())

	//time.Sleep(2 * time.Second)

	s.ContractInstance = instance
	s.ContractAddress = address

	f := []byte(address.String())
	err = os.WriteFile("./sample.address", f, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Sample) GetContractInstance(contractAddress common.Address) *sample.Contracts {
	instance, err := sample.NewContracts(contractAddress, s.User.Cli)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

func (s *Sample) GetVersion() {
	v, err := s.ContractInstance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("sample contract version: %v", v))
}

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

func (s *Sample) GetItems(name string) {
	key := [32]byte{}
	copy(key[:], name)

	result, err := s.ContractInstance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("item value: ", string(result[:]))
}
