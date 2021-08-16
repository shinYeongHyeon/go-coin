package main

import (
	"github.com/shinYeongHyeon/go-coin/cli"
	"github.com/shinYeongHyeon/go-coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
