package account

import (
	"context"
	"crypto/ecdsa"
	"evmosInteractor/logger"
	"evmosInteractor/utils"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"os/exec"
)

var User *UserStruct

type UserStruct struct {
	Cli  *ethclient.Client
	Pri  *ecdsa.PrivateKey
	Addr common.Address
}

// NewUser gets the private key of mykey account from evmos, and sets the user.
func NewUser() *UserStruct {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	data, err := exec.Command("evmosd", "keys", "unsafe-export-eth-key", "mykey", "--keyring-backend", "test").Output()
	if err != nil {
		log.Fatal(err)
	}

	privateKey, address := GetPriAndAddr(string(data)[:len(string(data))-1])

	return &UserStruct{
		Cli:  client,
		Pri:  privateKey,
		Addr: address,
	}
}

// GetBalance gets the balance of the account.
func (u *UserStruct) GetBalance() *big.Float {
	balance, err := u.Cli.BalanceAt(context.Background(), u.Addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue
}

// GetBalanceOf returns balance from a given account
func (u *UserStruct) GetBalanceOf(addr string) *big.Float {
	balance, err := u.Cli.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue
}

// GetNonce gets the nonce of the account, each transaction should be different.
func (u *UserStruct) GetNonce() uint64 {
	nonce, err := u.Cli.PendingNonceAt(context.Background(), u.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return nonce
}

// GetGasPrice returns suggest gas price.
func (u *UserStruct) GetGasPrice() *big.Int {
	gasPrice, err := u.Cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return gasPrice
}

// GetAuth sets the details of the transaction.
func (u *UserStruct) GetAuth() *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(u.Pri, big.NewInt(9000))
	if err != nil {
		fmt.Println("err tx: ", err)
	}
	auth.Nonce = big.NewInt(int64(u.GetNonce()))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = u.GetGasPrice()

	return auth
}

// VerifyPrivateKey verifies if the inserted private key and address is matched.
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

// GetPriAndAddr returns the private key in *ecdsa.PrivateKy form and address from the given private key.
func GetPriAndAddr(pri string) (*ecdsa.PrivateKey, common.Address) {
	private, err := crypto.HexToECDSA(pri)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := private.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return private, crypto.PubkeyToAddress(*publicKeyECDSA)
}

// SendEVMOS sends evmos from mykey to another account
func (u *UserStruct) SendEVMOS(to string, amount string) {
	value := utils.GetBigInt(amount)
	addr := common.HexToAddress(to)

	chainID, err := u.Cli.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    u.GetNonce(),
		GasPrice: u.GetGasPrice(),
		Gas:      300000,
		To:       &addr,
		Value:    value,
		Data:     common.Hex2Bytes("0x"),
		V:        nil,
		R:        nil,
		S:        nil,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), u.Pri)
	if err != nil {
		log.Fatal(err)
	}

	err = u.Cli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("send tx hash: ", signedTx.Hash().Hex())

	txData := logger.Transaction{
		From:   u.Addr.Hex(),
		To:     to,
		TxHash: signedTx.Hash().Hex(),
		Time:   time.Now().UTC(),
	}

	txData.UploadIPFS("eth", "transfer")
}
