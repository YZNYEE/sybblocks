package mongoDB

import (
	"SYNBLOCK/entity"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	_ "math/big"
)

var url string = "mongodb://localhost:27017"
var clientoptions *options.ClientOptions
var client *mongo.Client
var DB *mongo.Database
var err error
var Coll *mongo.Collection

func InsertBlock(mongoblock *entity.Mongoblock, collection *mongo.Collection) {
	fmt.Println(mongoblock)
	collection.InsertOne(context.TODO(), mongoblock)
}

func FindBlockbyNum(n uint64) *entity.Mongoblock {
	var b entity.Mongoblock
	q := bson.D{{"id", n}}
	//q := MongoBlock{Id: n}
	err := Coll.FindOne(context.TODO(), q).Decode(&b)
	if err != nil {
		log.Fatal(err)
	}
	return &b

}

type base struct {
	LastNum     uint64
	BucketsNum  uint64
	BucketsNums []uint64
}

func InitDB(lastnum, bn uint64) {
	b := base{lastnum, bn, make([]uint64, bn)}
	DB.Collection("base").InsertOne(context.TODO(), b)
}

func GetBase() base {
	var b base
	err := DB.Collection("base").FindOne(context.TODO(), bson.D{}).Decode(&b)
	if err != nil {
		InitDB(uint64(0), uint64(20))
		return base{uint64(0), uint64(20), make([]uint64, 20)}
	}
	return b
}

func GetLast() entity.Mongounit {
	var res entity.Mongounit
	Coll.FindOne(context.Background(), bson.D{{"key", "lastone"}}).Decode(&res)
	return res
}

func Insertlastone(lastone uint64) {
	var m entity.Mongounit
	m.Key = "lastone"
	m.Value = lastone
	Coll.InsertOne(context.TODO(), m)
}

func init() {
	clientoptions = options.Client().ApplyURI(url)
	client, err = mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database("test")
	Coll = DB.Collection("block")

}
