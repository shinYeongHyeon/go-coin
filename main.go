package main

import (
	"encoding/json"
	"fmt"
	"github.com/shinYeongHyeon/go-coin/blockchain"
	"github.com/shinYeongHyeon/go-coin/utils"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("http://localhost%s%s", port, u)), nil
}

type URLDescription struct {
	URL 		URL	   `json:"url"`
	Method 		string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: 		 URL("/"),
			Method: 	 "GET",
			Description: "See Documentation",
		},
		{
			URL: 		 URL("/blocks"),
			Method: 	 "POST",
			Description: "Add A Block",
			Payload:	 "data: string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	utils.HandleError(json.NewEncoder(rw).Encode(data))
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockChain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleError(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockChain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Println("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}