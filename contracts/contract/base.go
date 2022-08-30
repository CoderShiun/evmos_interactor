package contract

type IContract interface {
	Deploy()
}

func DeployContract(contract IContract) {
	contract.Deploy()
}
