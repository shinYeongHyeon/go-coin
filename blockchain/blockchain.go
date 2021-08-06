package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height	 int    `json:"height"`
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain // 변수의 인스턴스를 직접 공유하지 않음 -> Singleton Pattern
var once sync.Once

func (b *Block) calculateHash() {
	b.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.Data+ b.PrevHash)))
}

func getLastHash() string {
	totalBlocks := len(GetBlockChain().blocks)
	if totalBlocks == 0 {
		return ""
	}

	return GetBlockChain().blocks[totalBlocks - 1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block {data, "", getLastHash(), len(GetBlockChain().blocks) + 1}
	newBlock.calculateHash()

	return &newBlock
}

// AddBlock Add Block to chain
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

// GetBlockChain Getter BlockChain Instance
func GetBlockChain() *blockchain {
	if b == nil {
		// 단 한번만 실행하도록 도와주는 sync 라이브러리
		once.Do(func() {
			b = &blockchain {}
			b.AddBlock("Genesis Block")
		})
	}

	return b
}

// AllBlocks Get All Blocks
func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}