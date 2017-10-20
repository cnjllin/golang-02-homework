// Package  main provides ...
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["user"]

	if !ok || len(keys) < 1 {
		fmt.Fprintf(w, "Url Param 'user' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	fmt.Fprintf(w, "hello"+string(key))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8800", nil)
}
