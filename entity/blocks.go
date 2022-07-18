package entity

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Mongoblock struct {
	Id          uint64              `bson:"id"`
	BlockHash   string              `bson:"blockHash"`
	Transaction []*MongoTransaction `bson:"transaction"`
	Header      *types.Header       `bson:"header"`
	ParentHash  string              `bson:"parentHash"`
	Uncles      []*types.Header     `bson:"uncles"`
	Diffculty   uint64              `bson:"diffculty"`
	Bloom       types.Bloom         `bson:"bloom"`
	Gasused     uint64              `bson:"gasused"`
	Gaslimit    uint64              `bson:"gaslimit"`
	Time        uint64              `bson:"time"`
}
