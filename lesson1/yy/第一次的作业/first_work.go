package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello %s!</h1>", r.URL)
}

func user_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>user %s!</h1>", r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user/", user_handler)
	http.HandleFunc("/pc/", user_handler)
	http.ListenAndServe(":80", nil)
}
