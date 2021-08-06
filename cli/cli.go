package cli

import (
	"flag"
	"fmt"
	"github.com/shinYeongHyeon/go-coin/explorer"
	"github.com/shinYeongHyeon/go-coin/rest"
	"os"
)

func usage() {
	fmt.Printf("Welcome to GO-CLI-Playground\n\n")
	fmt.Printf("Please use the following flags\n\n")
	fmt.Printf("-port=4000 : Set the PORT of the server\n")
	fmt.Printf("-mode=rest: Choose between 'html' and 'rest'\n")
	os.Exit(1)
}

func Start() {
	if len(os.Args) < 2 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the Server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
