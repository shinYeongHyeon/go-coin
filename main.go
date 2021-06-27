package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello from home")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening On http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}