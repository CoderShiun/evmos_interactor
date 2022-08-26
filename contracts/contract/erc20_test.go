package contract

import (
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

	u := account.User{
		Cli:  client,
		Pri:  privateKey,
		Addr: address,
	}

	e.User = u

	e.DeployContract()

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
	Convey("check if erc20 contract instance is existed", t, func() {
		So(e.ContractInstance, ShouldNotBeEmpty)

		Convey("get erc20 token decimal, should be 18", func() {
			e.User.GetAuth()
			decimal, err := e.ContractInstance.Decimals(nil)
			if err != nil {
				t.Error(err)
			}

			So(decimal, ShouldEqual, uint8(18))
		})
	})
}
