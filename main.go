package main

import (
	"fmt"
	"net/http"

	"log"
	"github.com/WeiTing1991/mywebsite/view/index"
	"github.com/a-h/templ"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	indexHandler := index.Base("wellcome to my website")
	http.Handle("/", templ.Handler(indexHandler))
	log.Println("Server is running on port 8080")

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
