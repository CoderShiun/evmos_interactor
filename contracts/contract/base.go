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
