package main

import (
	"fmt"
	"net/http"
)

/*
http://localhost:8800?user=pc
响应一个 hello pc
*/
func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递过来的参数
	value := r.Form.Get("user")
	if value == "" || value != "pc" {
		fmt.Fprintf(w, "输入URL不符合要求，请重新输入此:http://localhost:8800/?user=pc\n")
	} else {
		fmt.Fprintf(w, "hello %s\n", value)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8800", nil)
}
