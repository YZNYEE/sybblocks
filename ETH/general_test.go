package ETH

import (
	"fmt"
	"testing"
)

func TestGetblocknum(t *testing.T) {
	n := Getblocknum()
	fmt.Println(n)

}
func TestGetBlockbyNum(t *testing.T) {
	//b2 := GetBlockbyNum(1)
	b1 := GetBlockbyNum(12594202)

	fmt.Println(b1.Transactions(), len(b1.Transactions()))

	//fmt.Println(b1.Header().)
	//fmt.Println(b2.Header().ParentHash.Hex())
	//fmt.Println(b.Number().Uint64())     // 5671744
	//fmt.Println(b.Time())                // 1527211625
	//fmt.Println(b.Difficulty().Uint64()) // 3217000136609065
	//fmt.Println(b.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	//fmt.Println(len(b.Transactions()))   // 144

	//count, err := client.TransactionCount(context.Background(), b.Hash())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(count) // 144

}
