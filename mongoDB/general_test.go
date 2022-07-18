package mongoDB

import (
	"SYNBLOCK/ETH"
	"SYNBLOCK/utils"
	"fmt"
	"log"
	"testing"
)

func TestInsertBlock(t *testing.T) {
	b, err := ETH.GetBlockbyNum(12594194)
	if err != nil {
		log.Fatal(err)
	}
	InsertBlock(utils.ParseBlock(b), Coll)
}

func TestFindBlockbyNum(t *testing.T) {
	b1, err := ETH.GetBlockbyNum(12594209)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b1.Hash().Hex())
	b := FindBlockbyNum(12594207)
	fmt.Println(b.BlockHash)
}

func TestInsertlastone(t *testing.T) {
	Insertlastone(10000)
	m := GetLast()
	fmt.Println(m)
}
