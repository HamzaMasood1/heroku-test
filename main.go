package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.HandleFunc("/test", TestingHandler)

	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":"+port, nil)
}

func TestingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello"))
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	fmt.Println(r.FormValue("body"))
	rw.Write(markdown)
}
