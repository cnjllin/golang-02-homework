package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "hello Golang", r.URL.Path)
	fmt.Fprintf(w, "hello Golang") // 如果添加r.URL.Path 请求的时候会提示 %!(EXTRA string=/) 对输出的格式没有做格式化
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
