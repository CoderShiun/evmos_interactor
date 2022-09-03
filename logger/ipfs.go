package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
)

var Sh = shell.NewShell("localhost:5001")

type Transaction struct {
	ContractAddress string    `json:"contract_address"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	TxHash          string    `json:"tx_hash"`
	Time            time.Time `json:"time"`
}

func (t *Transaction) UploadIPFS(txType string, function string) *shell.FilesStatObject {
	ctx := context.Background()

	fileName := fmt.Sprintf("/evmosInter/%v/%v-%v", txType, function, time.Now().UTC().String())

	err := Sh.FilesWrite(
		ctx,
		fileName,
		bytes.NewBufferString(string(marshalStruct(*t))),
		func(builder *shell.RequestBuilder) error {
			builder.Option("create", true)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	file, err := Sh.FilesStat(ctx, fileName)
	if err != nil {
		log.Error(err)
	}

	fmt.Println("IPFS - saving done. Cid: ", file.Hash)

	return file
}

func marshalStruct(transaction Transaction) []byte {
	data, err := json.Marshal(&transaction)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func unmarshalStruct(str []byte) Transaction {
	var transaction Transaction
	err := json.Unmarshal(str, &transaction)
	if err != nil {
		log.Fatal(err)
	}

	return transaction
}
