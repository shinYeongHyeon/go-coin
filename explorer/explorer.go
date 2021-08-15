package explorer

import (
	"fmt"
	"github.com/shinYeongHyeon/go-coin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const templateDir string = "explorer/templates/"

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home (writer http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(writer, "home", homeData {
		PageTitle: "Home",
		Blocks: nil,//Blocks: blockchain.GetBlockChain().AllBlocks(),
	})
}

func add (writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		templates.ExecuteTemplate(writer, "add", nil)
	case "POST":
		request.ParseForm()
		data := request.Form.Get("blockData")
		blockchain.BlockChain().AddBlock(data)
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}
}

var templates *template.Template

func Start(aPort int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening On http://localhost%s\n", fmt.Sprintf(":%d", aPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", aPort), handler))
}