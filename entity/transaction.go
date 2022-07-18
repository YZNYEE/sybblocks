package entity

import (
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
