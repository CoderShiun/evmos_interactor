package keys

import (
	"crypto/ecdsa"
	"evmosInteractor/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetNewAccount returns a new account.
func GetNewAccount() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("private: ", privateKey)

	// convert to byte array
	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println("private Byte: ", privateKeyBytes)

	// remove 0x
	//fmt.Println("private remove 0x: ", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	//fmt.Println("public: ", publicKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//fmt.Println("publicKeyECDSA: ", publicKey)

	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println("publicBytes to string: ", string(publicKeyBytes))
	// remove 0x04
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// from public key to address
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

// KeyStore is a keystore which can be used to store and retrieve key pairs.
func KeyStore(password string) {
	basicPath, err := utils.HomePath()
	if err != nil {
		log.Fatal(err)
	}

	ks := keystore.NewKeyStore(fmt.Sprintf("%v/.evmosd/keyring-test", basicPath), keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

// GetKeyFromStore reads a key from a file.
func GetKeyFromStore() {
	basicPath, err := utils.HomePath()
	if err != nil {
		log.Fatal(err)
	}
	file := fmt.Sprintf("%v/.evmosd/keyring-test/mykey.info", basicPath)

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := ""
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	if err = os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
