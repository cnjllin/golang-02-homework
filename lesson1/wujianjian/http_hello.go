package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request) {
	//fmt.Fprintf(w,"Hello,world!")
	fmt.Fprintf(w,"<h1>hello,%s!</h1>",r.URL.Path)
	fmt.Fprintf(w,"hello,%s!",r.RemoteAddr)
}
func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}
