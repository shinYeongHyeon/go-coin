package main

import "github.com/shinYeongHyeon/go-coin/blockchain"

func main() {
	blockchain.BlockChain().AddBlock("First")
	blockchain.BlockChain().AddBlock("Second")
	blockchain.BlockChain().AddBlock("Third")
}
