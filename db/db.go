package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/shinYeongHyeon/go-coin/utils"
)

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleError(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleError(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))

			return err
		})
		utils.HandleError(err)
	}

	return db
}

func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saveing Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleError(err)
}

func SaveBlockChain(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte("checkpoint"), data)
		return err
	})
	utils.HandleError(err)
}