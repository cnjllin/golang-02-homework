package main

import (
	"fmt"
	"net/http"
	"net/url"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/favicon.ico" {
		// chrome 每次都会请求 /favicon.ico
		return
	}
	//fmt.Fprintf(w, "<h1>hello, %s</h1>", r.URL.Path)
	// http://localhost:8800?user=pc
	// 响应一个 hello pc
	//fmt.Fprintf(w, "hello, %s", r.URL.RawQuery)
	fmt.Printf("URL: %s \n", r.URL)
	fmt.Printf("RawPath: %s \n", r.URL.RawPath)
	fmt.Printf("Path: %s \n", r.URL.Path)
	fmt.Printf("Scheme: %s \n", r.URL.Scheme)
	fmt.Printf("RawQuery: %s \n", r.URL.RawQuery)
	fmt.Printf("Fragment: %s \n", r.URL.Fragment)
	fmt.Printf("Opaque: %s \n", r.URL.Opaque)

	m, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		fmt.Fprintf(w, "ERROR")
		log.Fatal(err)
	} else {
		fmt.Printf("m: %s, user: %s \n", m, m["user"][0])
		fmt.Fprintf(w, "<h1>Hello, %s!</h1>", m["user"][0])
	}
}

func user_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>User: %s</h1>", r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user/", user_handler)
	http.ListenAndServe(":8800", nil)
}
