package entity

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type MongoTransaction struct {
	BlockNum  *big.Int `bson:"blockNum"`
	Blockhash string   `bson:"blockhash"`
	From      string   `bson:"from"`
	To        string   `bson:"to"`
	GasUsed   *big.Int `bson:"gasUsed"`
	Value     *big.Int `bson:"value"`
	Nouce     uint64   `bson:"nouce"`
	Data      []byte   `bson:"data"`
}

func ParseFromtypestx(transaction *types.Transaction, block *types.Block) *MongoTransaction {
	var m MongoTransaction
	to := transaction.To().Hex()
	singer := types.NewEIP2930Signer(transaction.ChainId())
	_from, err := singer.Sender(transaction)
	if err != nil {
		fmt.Println("from解析错误:", transaction.Hash().Hex())
		return nil
	}
	from := _from.Hex()
	gaslimit := big.NewInt(int64(transaction.Gas()))
	nouce := transaction.Nonce()
	data := transaction.Data()
	m.BlockNum = block.Number()
	m.Blockhash = block.Hash().Hex()
	m.From = from
	m.To = to
	m.Value = transaction.Value()
	m.Nouce = nouce
	m.Data = data
	m.GasUsed = gaslimit
	return &m

}
