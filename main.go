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
	router.GET("/", Index)

	router.GET("/hello/:name", Hello)

	//http.HandleFunc("/markdown", GenerateMarkdown)
	//http.HandleFunc("/test", TestingHandler)
	//
	//http.Handle("/", http.FileServer(http.Dir("public")))

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

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := fmt.Fprintln(writer, "Welcome!")
	if err != nil {
		return
	}
}

func TestingHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello"))
	if err != nil {
		return
	}
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	fmt.Println(r.FormValue("body"))
	_, err := rw.Write(markdown)
	if err != nil {
		return
	}
}
