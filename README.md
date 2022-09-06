# evmos_interactor

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

A CLI tool that can deploy smart contracts and call smart contract functions on the [Evmos](https://docs.evmos.org/) chain. It saves all the transaction hash in [IPFS mutable file system(MFS)](https://docs.ipfs.tech/concepts/file-systems/#mutable-file-system-mfs). With IPFS web UI, you are able to fine all your transaction history easily.
![image](https://github.com/CoderShiun/evmos_interactor/blob/main/img/ipfs01.png)
![image](https://github.com/CoderShiun/evmos_interactor/blob/main/img/ifps02.png)

This repository contains two default smart contracts:
1. [Sample smart contract](https://github.com/CoderShiun/evmos_interactor/blob/main/contracts/sample/Sample.sol) is a simple version samrt contract.
2. [ERC20 smart contract](https://github.com/CoderShiun/evmos_interactor/blob/main/contracts/erc20/ERC20.sol) is a smart contract which implements [EIP-20 token standard](https://eips.ethereum.org/EIPS/eip-20).

## Table of Contents
- [Background](#background)
- [Requirements](#requirements)
- [Install](#install)
- [Usage](#usage)
- [Tests](#tests)
- [Afterwards](#afterwards)
- [Maintainers](#maintainers)
- [License](#license)

## Background
evmos_interactor is a command line tool for using the standard go-ether RPC requests on the EVM-compatible chain, for example, evmos. It also includes two default smart contracts that can be used to deploy and interact with the contract, such as burn, mint and transfer tokens.

The goals for this repository are:
1. Compile and deploy an **ERC20** token smart contract to the local evmos test network.
2. Using **Go** to programmatically deploy the contract.
3. **Query and transfer token balances** on the deployed smart contract.
4. Test Driven Development **(TDD)**.
5. Behaviour Driven Development **(BDD)** approach to test expected behaviour.
6. Integration/**E2E** tests that use the JSON-RPC client to sent transactions to the node.

## Requirements
- [Go v1.18+](https://go.dev/)

## Install
This project simply use Go to install the tool.
Under the clone folder:
```sh
$ go mod tidy && go install
```
If you do not want to install it, use the fallowing cammand instead.
```sh
$ go run main.go
```

## Usage
**Please ensure your evmos local evmos test network is running in --keyring-backend=test mode, and mykey.info file is under the keyring-test folder.**

Shows the balance of mykey account based on the evmos chain.
```sh
$ evmosInterator account balance
```
Shows the default contract list.
```sh
$ evmosInterator contract list
```
Deploy the simply smart contract on the local evmos test network, and save contract address in the current folder.
```sh
$ evmosInterator contract deploy simply
```
Get the sample contract version.
```sh
$ evmosInterator contract sample version
```
Insert an item to sample contract.
```sh
$ evmosInterator contract sample set_item
```
Get item value from sample contract.
```sh
$ evmosInterator contract sample get_item
```
Deploy the erc20 smart contract on the local evmos test network, and save the contract address in current folder.
```sh
$ evmosInterator contract deploy erc20
```
Show erc20 token total supply.
```sh
$ evmosInterator contract erc20 total_supply
```
Show erc20 smart contract token decimal.
```sh
$ evmosInterator contract erc20 decimals
```
Get erc20 token balance.
```sh
$ evmosInterator contract erc20 get_balance
```
Mint amount of erc20 tokens to msg.sender.
```sh
$ evmosInterator contract erc20 mint
```
Burn an amount of erc20 tokens from the sender account.
```sh
$ evmosInterator contract erc20 burn
```
Transfer amount of erc20 tokens to another account.
```sh
$ evmosInterator contract erc20 transfer
```

## Tests
TDD and BDD tests: because it is pressed for time, there are few tests in account_test.go and erc20_test.go. BDD test uses [Goconvey](https://github.com/smartystreets/goconvey). 
```sh
$ go run test -v
```
E2E test - an auto test progress under the evmosInterator/E2E/interE2E. Make sure you have already installed the evmosInteractor tool first. 
Under the test folder, run:
```sh
$ go install
```
Execute the program to see the results:
```sh
$ interE2E 
```

## Afterwards

The first idea was just to write a simple program to run through a few processes like contract deployment, send tx and so on. But after reading the codes from evmos, it seems like a good idea to build a command line tool as well. The plan was to use docker to ensure everyone has the same environment, but because of the limit of the time and also need to get familiar with evmos chain and tool, I had to drop the docker.

## Maintainers
[@CoderShiun](https://github.com/CoderShiun).

## License
Apache-2.0 license
