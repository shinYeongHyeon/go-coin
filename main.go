package main

import (
	"fmt"
	"github.com/shinYeongHyeon/go-coin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const (
	port 		string = ":4000"
	templateDir string = "templates/"

)
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home (writer http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(writer, "home", homeData {
		PageTitle: "Home",
		Blocks: blockchain.GetBlockChain().AllBlocks(),
	})
}

func add (writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		templates.ExecuteTemplate(writer, "add", nil)
	case "POST":
		request.ParseForm()
		data := request.Form.Get("blockData")
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening On http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}