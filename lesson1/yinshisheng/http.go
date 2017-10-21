package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>hi,51reboot!</p>")
}

func main() {
	http.HandleFunc("/", handler)  //目录可以使用正则匹配吗？
	http.ListenAndServe(":8080", nil) //发现必须加冒号,不加编译也不报错，但是不监听端口。
}
