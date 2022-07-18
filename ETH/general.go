package ETH

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)
import "context"

func Getblocknum() uint64 {
	num, errn := client.BlockNumber(context.Background())
	if errn != nil {
		log.Fatal(errn)
	}
	return num
}

func GetBlockbyNum(num uint64) (*types.Block, error) {
	bignum := big.NewInt(int64(num))
	//client.SyncProgress()
	block, errb := client.BlockByNumber(context.Background(), bignum)
	if errb != nil {
		return nil, errb
	}
	return block, err

}

//func Syncblocks(block *types.Block) {
//	b := entity.ParseBlock(block)
//	mongoDB.InsertBlock(b)
//}

var client *ethclient.Client
var err error

func init() {
	client, err = ethclient.Dial("https://ropsten.infura.io/v3/4065fc77233f4f18886d7236e9042ca6")

}
