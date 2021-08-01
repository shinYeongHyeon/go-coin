package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}