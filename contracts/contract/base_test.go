package contract

import (
	"evmosInteractor/contracts/account"
	"os/exec"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

var s Sample

func TestDeploy(t *testing.T) {
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

	// Erc20 contract
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

	// Sample contract
	s.User = u
	s.Deploy()

	cmd = exec.Command("rm", "sample.address")
	err = cmd.Run()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(2 * time.Second)

	if s.ContractAddress.String() == "" {
		t.Error("no contract address read")
	}

	if s.ContractInstance == nil {
		t.Error("no contract instance set")
	}
}
