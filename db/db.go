package db

import (
	"github.com/boltdb/bolt"
	"github.com/shinYeongHyeon/go-coin/utils"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPoint, err := bolt.Open("blockchain.db", 0600, nil)
		utils.HandleError(err)
		db = dbPoint
	}

	return db
}
