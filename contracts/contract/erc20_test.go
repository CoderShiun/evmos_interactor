package contract

import (
	"evmosInteractor/keys"
	"github.com/ethereum/go-ethereum/common"
	"os/exec"
	"testing"
	"time"

	"evmosInteractor/contracts/account"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/smartystreets/goconvey/convey"
)

var e ERC20

func TestDeployContract(t *testing.T) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		t.Fatal(err)
	}

	data, err := exec.Command("evmosd", "keys", "unsafe-export-eth-key", "mykey", "--keyring-backend", "test").Output()
	if err != nil {
		t.Fatal(err)
	}

	privateKey, address := account.GetPriAndAddr(string(data)[:len(string(data))-1])

	u := account.UserStruct{
		Cli:  client,
		Pri:  privateKey,
		Addr: address,
	}

	e.User = u

	e.Deploy()

	cmd := exec.Command("rm", "erc20.address")
	err = cmd.Run()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(2 * time.Second)

	if e.ContractAddress.String() == "" {
		t.Error("no contract address read")
	}

	if e.ContractInstance == nil {
		t.Error("no contract instance set")
	}
}

func TestGetDecimals(t *testing.T) {
	Convey("GetDecimals", t, func() {
		Convey("Check if erc20 contract instance is existed", func() {
			So(e.ContractInstance, ShouldNotBeEmpty)

			Convey("Get erc20 token decimal, should be 18", func() {
				decimal := e.GetDecimals()

				So(decimal, ShouldEqual, uint8(18))
			})
		})
	})
}

func TestGetTotalSupply(t *testing.T) {
	Convey("GetTotalSupply", t, func() {
		Convey("erc20 initial total supply should be 1000000000000000000000", func() {
			total := e.GetTotalSupply().String()

			So(total, ShouldEqual, "1000000000000000000000")
		})
	})
}

func TestBalanceOf(t *testing.T) {
	Convey("BalanceOf", t, func() {
		Convey("mykey initial balance should be the same amount of total supply", func() {
			total, err := e.ContractInstance.TotalSupply(nil)
			if err != nil {
				t.Fatal(err)
			}

			balance := e.BalanceOf(e.User.Addr).Int64()

			So(total.Int64(), ShouldEqual, balance)
		})
	})
}

func TestMint(t *testing.T) {
	Convey("Mint", t, func() {
		Convey("Mint 2000 tokens, then check if total supply is correct", func() {
			e.Mint("2000")
			time.Sleep(2 * time.Second)

			total := e.GetTotalSupply().String()
			So(total, ShouldEqual, "1000000000000000002000")
		})
	})
}

func TestBurn(t *testing.T) {
	Convey("Burn", t, func() {
		Convey("Burn 1000 tokens, then check if total supply is correct", func() {
			e.Burn("1000")
			time.Sleep(2 * time.Second)

			total := e.GetTotalSupply().String()
			So(total, ShouldEqual, "1000000000000000001000")
		})
	})
}

func TestTransfer(t *testing.T) {
	Convey("Transfer", t, func() {
		Convey("Generate a new random address", func() {
			newAddr := keys.GetNewAccount()
			So(newAddr, ShouldNotEqual, e.User.Addr)

			Convey("Send 1,000 tokens to new address", func() {
				e.Transfer(newAddr, "1000")
				time.Sleep(2 * time.Second)

				balance, err := e.ContractInstance.BalanceOf(nil, common.HexToAddress(newAddr))
				if err != nil {
					t.Fatal(err)
				}

				So(balance.Int64(), ShouldEqual, 1000)
			})
		})
	})
}
