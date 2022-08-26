package keys

import (
	"crypto/ecdsa"
	"evmosInteract/utils"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"os"
)

func GetNewKey2(data []byte) {
	private2 := secp256k1.PrivKeyFromBytes(data)

	privateKey := private2.ToECDSA()

	// convert to byte array
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private Byte: ", privateKeyBytes)

	// remove 0x
	fmt.Println("private remove 0x: ", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	fmt.Println("public: ", publicKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fmt.Println("publicKeyECDSA: ", publicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicBytes to string: ", string(publicKeyBytes))
	// remove 0x04
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// from public key to address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)
}

func GetNewKey() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("private: ", privateKey)

	// convert to byte array
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private Byte: ", privateKeyBytes)

	// remove 0x
	fmt.Println("private remove 0x: ", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	fmt.Println("public: ", publicKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fmt.Println("publicKeyECDSA: ", publicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicBytes to string: ", string(publicKeyBytes))
	// remove 0x04
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// from public key to address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)
}

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

func rest(dat []byte) {
	kh := crypto.Keccak256Hash(dat)
	kh.Hex()

	fmt.Println(hexutil.Encode(crypto.Keccak256(dat)))

	private, err := crypto.HexToECDSA(kh.Hex()[2:])
	if err != nil {
		log.Fatal(err)
	}

	//private, err := crypto.ToECDSA(dat)
	//private, err := crypto.LoadECDSA(dat)
	//fmt.Println(crypto.ToECDSA(crypto.Keccak256(dat)))

	publicKey := private.Public()
	//fmt.Println("public: ", publicKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//fmt.Println("publicKeyECDSA: ", publicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println("publicBytes to string: ", string(publicKeyBytes))
	// remove 0x04
	fmt.Println(hexutil.Encode(publicKeyBytes))

	// from public key to address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)
}
