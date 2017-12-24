package main

import (
	_ "./statik" // TODO: Replace with the absolute import path
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

// ...
func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}
