package main

import (
	"SYNBLOCK/ETH"
	"SYNBLOCK/mongoDB"
	"fmt"
)

var exit chan bool

func LoadBlockInfo(start, end uint64) {
	n := ETH.Getblocknum()
	fmt.Println(n)
	for i := uint64(start); i <= uint64(end); i++ {
		fmt.Println(i)
		b := ETH.GetBlockbyNum(i)
		out := fmt.Sprintf("loading##%d/%d##blocks,blocks hash:%s", i, n, b.Hash().Hex())
		fmt.Println(out)
		mongoDB.InsertBlock(b)
	}
	exit <- true
}

func main() {
	//works := 16
	exit = make(chan bool, 16)
	n := ETH.Getblocknum()
	fmt.Println(n)
	base := int(n / 16)
	for i := 0; i < 16; i++ {
		go LoadBlockInfo(uint64(base*i), uint64(base*(i+1)-1))
		//fmt.Println("range", base*i, base*(i+1)-1)
	}
	//go LoadBlockInfo(22258, 50000)
	//go LoadBlockInfo(50001, 100000)
	//go LoadBlockInfo(100001, 150000)
	//go LoadBlockInfo(150001, 200000)
	for i := 0; i < 16; i++ {
		<-exit
	}
}
