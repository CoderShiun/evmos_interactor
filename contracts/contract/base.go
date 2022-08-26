package contract

import (
	"evmosInteractor/contracts/erc20"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IContract interface {
	DeployContract()
	GetContractInstance(c *ethclient.Client, address common.Address) *erc20.Contracts
}

/*func DeployContract(c *ethclient.Client, address common.Address, privateKey *ecdsa.PrivateKey) common.Address {
	nonce, err := c.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("my nonce: ", nonce)

	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice: ", gasPrice)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(9000))
	if err != nil {
		fmt.Println("err tx: ", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := contracts.DeployContracts(auth, c, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance

	return address
}

func GetContractInstance(c *ethclient.Client, address common.Address) *contracts.Contracts {
	instance, err := contracts.NewContracts(address, c)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}*/
