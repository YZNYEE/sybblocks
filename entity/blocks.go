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

func ParseBlock(block *types.Block) *Mongoblock {
	txs := make([]*MongoTransaction, 0)
	for i := 0; i < len(block.Transactions()); i++ {
		if block.Transactions()[i].To() == nil {
			continue
		}
		tx := ParseFromtypestx(block.Transactions()[i], block)
		//fmt.Println("#######", i, tx)
		if tx != nil {
			break
		}
		txs = append(txs, tx)
	}
	var b Mongoblock
	b.BlockHash = block.Hash().Hex()
	b.Id = block.Number().Uint64()
	header := block.Header()
	b.Header = header
	b.ParentHash = block.ParentHash().Hex()
	b.Uncles = block.Uncles()
	b.Diffculty = block.Difficulty().Uint64()
	time := block.Time()
	b.Time = time
	gasused := block.GasUsed()
	b.Gasused = gasused
	bloom := block.Bloom()
	b.Bloom = bloom
	b.Gaslimit = block.GasLimit()
	return &b
}
