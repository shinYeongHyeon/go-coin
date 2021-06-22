package blockchain

type block struct {
	data string
	hash string
	prevHash string
}

type blockchain struct {
	blocks []block
}

var b *blockchain // 변수의 인스턴스를 직접 공유하지 않음 -> Singleton Pattern

// GetBlockChain Getter BlockChain Instance
func GetBlockChain() *blockchain {
	if b == nil {
		b = &blockchain {}
	}

	return b
}