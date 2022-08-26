# evmos_interactor
A CLI tool that is able to deploy smart contracts and call smart contract functions on the evmos chain.

There are two defalut smart contracts, sample.sol and erc20.sol.
You should first make sure your evmos local test node is running, then deploy the smart contract on evmos chan, 
e.g. $go run main.go contract deploy {contract_name}.
Can check contract name by $go run main.go contract list.

After the deployment, it saves the contract address on local.
Call contract function please use $go run main.go contract {contract_name} and follow the guide.
