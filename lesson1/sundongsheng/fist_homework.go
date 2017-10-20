package main

import (
	"net/http"
	"fmt"
)

func get_handler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Fprintf(w,"hello %s", r.Form.Get("user"))
}

func main() {
	http.HandleFunc("/", get_handler)
	http.ListenAndServe(":8800", nil)
}
