package main

import (
	"github.com/shinYeongHyeon/go-coin/explorer"
	"github.com/shinYeongHyeon/go-coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}