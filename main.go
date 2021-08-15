package main

import (
	"github.com/shinYeongHyeon/go-coin/blockchain"
	"github.com/shinYeongHyeon/go-coin/cli"
)

func main() {
	blockchain.BlockChain()
	cli.Start()
}
