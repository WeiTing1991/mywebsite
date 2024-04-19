package main

import (
	"net/http"

	"os"
	"log"
	"github.com/WeiTing1991/mywebsite/view/index"
	"github.com/WeiTing1991/mywebsite/view/work"
	"github.com/a-h/templ"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	indexHandler := index.Base("weitingchen")
	http.Handle("/", templ.Handler(indexHandler))

	workHandler := work.Base("weitingwork")
	http.Handle("/work", templ.Handler(workHandler))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server is running on port 8080")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
