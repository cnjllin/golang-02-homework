package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, %s</h1>", r.URL.Path)
	fmt.Fprintf(w, "<h1>Hello, %s</h1>", r.URL)
	// 请求IP
	fmt.Fprintf(w, "<h1>Client IP: %s</h1>", r.RemoteAddr)
	//fmt.Fprintf(w, "<h1>Hello, %s</h1>", r.WriteProxy())
}

func useer_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>user %s</h1>", r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	// curl http://127.0.0.1:8080/user/adfd
	http.HandleFunc("/user/", useer_handler)
	http.ListenAndServe(":8080", nil)
}