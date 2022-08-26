package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"os/exec"
)

type User struct {
	Cli  *ethclient.Client
	Pri  *ecdsa.PrivateKey
	Addr common.Address
}

func NewUser() *User {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}
	data, err := exec.Command("evmosd", "keys", "unsafe-export-eth-key", "mykey", "--keyring-backend", "test").Output()
	if err != nil {
		log.Fatal(err)
	}

	addr := common.HexToAddress("0xc6dcbd80a50218ff44481babce8aa757b5413d32")

	privateKey, ok := VerifyPrivateKey(string(data)[:len(string(data))-1], addr)
	if !ok {
		log.Fatal("wrong private key")
	}

	return &User{
		Cli:  client,
		Pri:  privateKey,
		Addr: addr,
	}
}

func (u *User) GetBalance() *big.Float {
	//account := common.HexToAddress(address)
	balance, err := u.Cli.BalanceAt(context.Background(), u.Addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(20)))

	return ethValue
}

func (u *User) GetNonce() uint64 {
	nonce, err := u.Cli.PendingNonceAt(context.Background(), u.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return nonce
}

func (u *User) GetGasPrice() *big.Int {
	gasPrice, err := u.Cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return gasPrice
}

func (u *User) GetAuth() *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(u.Pri, big.NewInt(9000))
	if err != nil {
		fmt.Println("err tx: ", err)
	}
	auth.Nonce = big.NewInt(int64(u.GetNonce()))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = u.GetGasPrice()

	return auth
}

func VerifyPrivateKey(pri string, addr common.Address) (*ecdsa.PrivateKey, bool) {
	private, err := crypto.HexToECDSA(pri)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := private.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println("address: ", address)

	if address == addr.String() {
		return private, true
	}
	return nil, false
}
