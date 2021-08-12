package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height	   int    `json:"height"`
}

var b *blockchain // 변수의 인스턴스를 직접 공유하지 않음 -> Singleton Pattern
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

// BlockChain Getter BlockChain Instance
func BlockChain() *blockchain {
	if b == nil {
		// 단 한번만 실행하도록 도와주는 sync 라이브러리
		once.Do(func() {
			b = &blockchain {"", 0}
			b.AddBlock("Genesis Block")
		})
	}

	return b
}