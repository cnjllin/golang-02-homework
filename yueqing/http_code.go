package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Golang", r.URL.Path)
}

// 仿照写一个函数
func handler_f(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "这是一个测试的函数", r.URL.Path)
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", handler_f)
	http.ListenAndServe(":20080", nil)
}
