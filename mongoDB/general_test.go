package mongoDB

import (
	"SYNBLOCK/ETH"
	"fmt"
	"testing"
)

//func TestInsertBlock(t *testing.T) {
//	b := ETH.GetBlockbyNum(12594203)
//	InsertBlock(entity.ParseBlock(b))
//}

func TestFindBlockbyNum(t *testing.T) {
	b1 := ETH.GetBlockbyNum(12594203)
	fmt.Println(b1.Hash().Hex())
	b := FindBlockbyNum(12594203)
	fmt.Println(b.BlockHash)
}

func TestInsertlastone(t *testing.T) {
	Insertlastone(10000)
	m := GetLast()
	fmt.Println(m)
}
