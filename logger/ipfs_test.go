package logger

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUploadIPFS(t *testing.T) {
	Convey("UploadIPFS", t, func() {
		Convey("Upload a new file to path /evmosInter/eth", func() {
			now := time.Now().UTC()

			tx := Transaction{
				ContractAddress: "0xf754763bd0d48a6a01760a6ea06ca85d71e258952bb73f3d175826657175cd15",
				From:            "0x332810ce3096c18c71f66ab8e468c75fb177631285b05bdb836d1af8e6197144",
				To:              "0xcb2990b374fd865ec61e0de02408096586c89f3336da903957a9d64aeb9eb693",
				TxHash:          "0x586bc03bb34a30a1479ea09b68bccaba3fb88224df080f1be78de85a9f6ec15a",
				Time:            now,
			}

			fileState := tx.UploadIPFS("eth", "test")
			So(fileState.Type, ShouldEqual, "file")
			So(fileState.Size, ShouldBeGreaterThan, 0)

			Convey("Get file back from IPFS", func() {
				//ctx := context.Background()

				time.Sleep(2 * time.Second)

				txByte := marshalStruct(tx)

				cat, err := Sh.Cat(fileState.Hash)
				if err != nil {
					t.Fatal(err)
				}
				defer cat.Close()

				content := make([]byte, len(txByte))
				time.Sleep(1 * time.Second)

				_, err = cat.Read(content)
				if err != nil {
					t.Fatal(err)
				}

				marshalTx := marshalStruct(tx)

				So(string(marshalTx), ShouldEqual, string(content))
			})
		})
	})
}
