package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/shinYeongHyeon/go-coin/db"
	"github.com/shinYeongHyeon/go-coin/utils"
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height	   int    `json:"height"`
}

var b *blockchain // 변수의 인스턴스를 직접 공유하지 않음 -> Singleton Pattern
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.HandleError(gob.NewDecoder(bytes.NewReader(data)).Decode(b))
}

func (b *blockchain) persist() {
	db.SaveBlockChain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// BlockChain Getter of BlockChain Instance
func BlockChain() *blockchain {
	if b == nil {
		// 단 한번만 실행하도록 도와주는 sync 라이브러리
		once.Do(func() {
			b = &blockchain {"", 0}
			fmt.Printf("NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				fmt.Println("Restoring...")
				b.restore(checkpoint)
			}
		})
	}

	fmt.Printf("NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
	return b
}