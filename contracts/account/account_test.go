package account

import (
	"crypto/ecdsa"
	"evmosInteractor/keys"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
)

var user = NewUser()

func TestVerifyPrivateKey(t *testing.T) {
	testCases := []struct {
		privateKey string
		address    common.Address
	}{
		{
			privateKey: "b08311615faa38f921d44f7661273ac54742dd013c64a9d3a5e21b7b62a132c9",
			address:    common.HexToAddress("0x4631c78ea953f8788E0Ae7262978eF9C2AeBaA0F"),
		},
		{
			privateKey: "30cb46431e4c1502702f39167dce14d174df4f1156e78cb7925d121a9578bb61",
			address:    common.HexToAddress("0xc6e7441D6B9F73d913254765E578eaE3aD224210"),
		},
	}

	for _, testCase := range testCases {
		pri, ok := VerifyPrivateKey(testCase.privateKey, testCase.address)
		t.Log("pri: ", hexutil.Encode(crypto.FromECDSA(pri))[2:])

		if hexutil.Encode(crypto.FromECDSA(pri))[2:] != testCase.privateKey || !ok {
			t.Errorf("VerifyPrivateKey failed, expect %s, got %s", testCase.privateKey, hexutil.Encode(crypto.FromECDSA(pri))[2:])
		}
	}
}

func TestGetPriAndAddr(t *testing.T) {
	Convey("GetPriAndAddr", t, func() {
		Convey("Given a new private key and address from GetPriAndAddr", func() {
			pri, addr := GetPriAndAddr("03c69f08909e9e75ef1b723d0a53bdcce946f04a9adcebc5c81b96c67e0587aa")

			Convey("When we convert them to a readable(string) value", func() {
				pri2 := hexutil.Encode(crypto.FromECDSA(pri))[2:]
				publicKey := pri.Public()
				publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
				addr2 := crypto.PubkeyToAddress(*publicKeyECDSA)

				Convey("Then the results should be the same", func() {
					So("03c69f08909e9e75ef1b723d0a53bdcce946f04a9adcebc5c81b96c67e0587aa", ShouldEqual, pri2)
					So(addr.Hex(), ShouldEqual, addr2.Hex())
				})
			})
		})

		Convey("should not return empty results", func() {
			pri, addr := GetPriAndAddr("6e4d2a625ed1aeaa8dbedc8d4b47c069406d3aaa4c61b44803f7158711be44b7")
			So(pri, ShouldNotBeEmpty)
			So(addr, ShouldNotBeEmpty)
		})
	})
}

func TestGetBalance(t *testing.T) {
	Convey("GetPriAndAddr", t, func() {
		Convey("Result should not be 0", func() {
			balance := user.GetBalance()
			So(balance, ShouldNotEqual, 0)
		})
	})
}

func TestGetBalanceOf(t *testing.T) {
	Convey("GetBalanceOf", t, func() {
		Convey("Get balance from mykey, amount should be the same of GetBalance function", func() {
			So(user.GetBalanceOf(user.Addr.Hex()).String(), ShouldEqual, user.GetBalance().String())
		})
	})
}

func TestSendEVMOS(t *testing.T) {
	Convey("SendEVMOS", t, func() {
		Convey("Given a new random address", func() {
			newAddr := keys.GetNewAccount()
			So(newAddr, ShouldNotEqual, user.Addr.Hex())

			Convey("When send token to new account, mykey balance should be reduced", func() {
				balance1, _ := user.GetBalance().Int64()
				user.SendEVMOS(newAddr, "1000000000000000000")
				time.Sleep(2 * time.Second)
				balance2, _ := user.GetBalance().Int64()
				So(balance1-1, ShouldEqual, balance2)

				Convey("Then get balance from new account, see if it receives correct amount of tokens", func() {
					balance, _ := user.GetBalanceOf(newAddr).Int64()
					So(balance, ShouldEqual, 1)
				})
			})
		})
	})
}
