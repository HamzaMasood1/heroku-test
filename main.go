package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := httprouter.New()

	router.GET("/hello/:name", Hello)
	router.GET("/test", TestingHandler)
	router.POST("/markdown", GenerateMarkdown)

	//serve static files from the root path "/"
	router.NotFound = http.FileServer(http.Dir("public"))

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return
	}
}

func Hello(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := fmt.Fprintf(writer, "Hello, %s!\n", params.ByName("name"))
	if err != nil {
		return
	}
}

func TestingHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := writer.Write([]byte("Hello"))
	if err != nil {
		return
	}
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	fmt.Println(r.FormValue("body"))
	_, err := rw.Write(markdown)
	if err != nil {
		return
	}
}
