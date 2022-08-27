A CLI tool that can deploy smart contracts and call smart contract functions on the evmos chain.

There are two default smart contracts, sample.sol and erc20.sol.  You should first ensure your evmos local evmos test network is running in --keyring-backend=test mode, and have mykey.info file under the keyring-test folder.

Use $go install to simply install the tool. If you do not want to install it, use $go run main.go instead.

Read the command introduction by using $evmosInterator

Commands:
1. $evmosInterator account balance - shows the balance of mykey account based on the evmos chain.
2. $evmosInterator contract list - shows the default contract list.
3. $evmosInterator contract deploy simply - deploy the simply smart contract on the local evmos test network, and save contract address in the current folder.
4. $evmosInterator contract sample version - get the sample contract version.
5. $evmosInterator contract sample set_item - insert an item to sample contract.
6. $evmosInterator contract sample get_item - get item value from sample contract.
7. $evmosInterator contract deploy erc20 - deploy the erc20 smart contract on the local evmos test network, and save the contract address in current folder.
8. $evmosInterator contract erc20 total_supply - show erc20 token total supply.
9. $evmosInterator contract erc20 decimals - show erc20 smart contract token decimal.
10. $evmosInterator contract erc20 get_balance - get erc20 token balance.
11. $evmosInterator contract erc20 mint - mint amount of erc20 tokens to msg.sender
12. $evmosInterator contract erc20 burn - burn an amount of erc20 tokens from the sender account.
13. $evmosInterator contract erc20 transfer - transfer amount of erc20 tokens to another account.

TDD and BDD tests: because it is pressed for time, I wrote a few tests in account_test.go and erc20_test.go. I use Goconvey for the BDD test. Please simply run $go run test -v

E2E test: I wrote an auto test progress under the evmosInterator/E2E/interE2E. Make sure you have already installed the evmosInteractor tool first. Run $go install under the test folder, and execute the program $interE2E to see the result.

Afterwards - The first idea was just to write a simple program to run through a few processes like contract deployment, send tx and so on. But after reading the codes from evmos, it seems like a good idea to build a command line tool as well. The plan was to use docker to ensure everyone has the same environment, but because of the limit of the time and also need to get familiar with evmos chain and tool, I had to drop the docker.
