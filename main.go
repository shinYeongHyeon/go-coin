package main

import (
	"fmt"
	"github.com/shinYeongHyeon/go-coin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home (writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	tmpl.Execute(writer, homeData {
		PageTitle: "Home",
		Blocks: blockchain.GetBlockChain().AllBlocks(),
	})
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening On http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}