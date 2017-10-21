package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello,world!", r.URL.Path)
	//      打印url 路径
	fmt.Fprintf(w, "Hello,word", r.RemoteAddr)
	// 访问127.0.0.1:8080 打印出ip 地址
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
