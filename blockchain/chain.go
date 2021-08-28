package blockchain

import (
	"fmt"
	"github.com/shinYeongHyeon/go-coin/db"
	"github.com/shinYeongHyeon/go-coin/utils"
	"sync"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval 	   int = 2
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height	   		  int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain // 변수의 인스턴스를 직접 공유하지 않음 -> Singleton Pattern
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}

	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval - 1]
	actualTime := (newestBlock.Timestamp/60) - (lastRecalculatedBlock.Timestamp/60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime < expectedTime {
		return b.CurrentDifficulty + 1
	} else if actualTime > expectedTime {
		return b.CurrentDifficulty - 1
	}

	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height & difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

// BlockChain Getter of BlockChain Instance
func BlockChain() *blockchain {
	if b == nil {
		// 단 한번만 실행하도록 도와주는 sync 라이브러리
		once.Do(func() {
			b = &blockchain {
				Height: 0,
			}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				b.restore(checkpoint)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}