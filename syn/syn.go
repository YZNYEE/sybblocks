package syn

import (
	"SYNBLOCK/ETH"
	"SYNBLOCK/entity"
	"fmt"
	"time"
)
import "SYNBLOCK/mongoDB"

func LoadBlockbyrange(start, end int) {
	buckets := start / 1000000
	for i := start; i <= end; i++ {
		b := entity.ParseBlock(ETH.GetBlockbyNum(uint64(i)))
		coll := mongoDB.DB.Collection(fmt.Sprint("buckets", buckets))
		mongoDB.InsertBlock(b, coll)
	}
}

//从当前存储同步区块
func Synblocks() {
	m := mongoDB.GetLast()
	lastone := m.Value.(uint64)
	for {
		lastone++
		block, err := ETH.GetBlockbyNum(lastone)
		if err != nil {
			fmt.Println(err)
			time.Sleep(15 * time.Second)
			continue
		}
		b := entity.ParseBlock(block)
		mongoDB.InsertBlock(b, mongoDB.DB.Collection("blocks"))
		mongoDB.Insertlastone(lastone)

	}
}
func Synfromzero() {
	//TODO
}
