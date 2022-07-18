package utils

import (
	"SYNBLOCK/entity"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func ParseBlock(block *types.Block) *entity.Mongoblock {
	txs := make([]*entity.MongoTransaction, 0)
	for i := 0; i < len(block.Transactions()); i++ {
		if block.Transactions()[i].To() == nil {
			continue
		}
		tx := ParseFromtypestx(block.Transactions()[i], block)
		//fmt.Println("#######", i, tx)
		if tx == nil {
			continue
		}
		txs = append(txs, tx)
	}
	var b entity.Mongoblock
	b.Transaction = txs
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

func ParseFromtypestx(transaction *types.Transaction, block *types.Block) *entity.MongoTransaction {
	var m entity.MongoTransaction
	to := transaction.To().Hex()
	singer := types.NewEIP2930Signer(transaction.ChainId())
	//fmt.Println(singer.ChainID(), "#######")
	_from, err := singer.Sender(transaction)
	if err != nil {
		//fmt.Println("！！！！解析错误:", err.Error())
		return nil
	} else {
		//fmt.Println("fff:", _from)
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
