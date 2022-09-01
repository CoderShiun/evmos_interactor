package account

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/smartystreets/goconvey/convey"
)

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
		Convey("GetPriAndAddr return private key and address", func() {
			pri, addr := getPriAndAddr("03c69f08909e9e75ef1b723d0a53bdcce946f04a9adcebc5c81b96c67e0587aa")

			Convey("When we convert them to a readable(string) value", func() {
				pri2 := hexutil.Encode(crypto.FromECDSA(pri))[2:]
				publicKey := pri.Public()
				publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
				addr2 := crypto.PubkeyToAddress(*publicKeyECDSA)

				Convey("Then the results should be the same.", func() {
					So("03c69f08909e9e75ef1b723d0a53bdcce946f04a9adcebc5c81b96c67e0587aa", ShouldEqual, pri2)
					So(addr.Hex(), ShouldEqual, addr2.Hex())
				})
			})
		})

		Convey("should not return empty results", func() {
			pri, addr := getPriAndAddr("6e4d2a625ed1aeaa8dbedc8d4b47c069406d3aaa4c61b44803f7158711be44b7")
			So(pri, ShouldNotBeEmpty)
			So(addr, ShouldNotBeEmpty)
		})
	})
}
