package main

import (
	"encoding/json"
	"fmt"
	"github.com/shinYeongHyeon/go-coin/utils"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL string
	Method string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: 		 "/",
			Method: 	 "GET",
			Description: "See Documentation",
		},
	}
	bytes, err := json.Marshal(data)
	utils.HandleError(err)
	rw.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s", bytes)

}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}