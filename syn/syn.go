package syn

import (
	"SYNBLOCK/ETH"
	"SYNBLOCK/utils"
	"fmt"
	"time"
)
import "SYNBLOCK/mongoDB"

//func LoadBlockbyrange(start, end int) {
//	buckets := start / 1000000
//	for i := start; i <= end; i++ {
//		b
//		b, err := entity.ParseBlock()
//		if err != nil {
//			log.Fatal(err)
//		}
//		coll := mongoDB.DB.Collection(fmt.Sprint("buckets", buckets))
//		mongoDB.InsertBlock(b, coll)
//	}
//}

//从当前存储同步区块
func Synblocks() {
	m := mongoDB.GetLast()
	lastone := m.Value.(float64)
	one := uint64(lastone)
	for {
		one++
		block, err := ETH.GetBlockbyNum(one)
		if err != nil {
			fmt.Println("加载:", one, "block出错", err)
			time.Sleep(15 * time.Second)
			continue
		}
		fmt.Println("加载:", one, "block成功")
		b := utils.ParseBlock(block)
		mongoDB.InsertBlock(b, mongoDB.DB.Collection("blocks"))
		mongoDB.Insertlastone(one)

	}
}
func Synfromzero() {
	//TODO
}
