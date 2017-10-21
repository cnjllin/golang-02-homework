package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	在http中，localhost:20080?user=pc 这个?user格式叫变量
*/
func handler(w http.ResponseWriter, r *http.Request) {
	// 获取参数 user 相应的值
	if ok := r.ParseForm(); ok != nil {
		log.Fatal(ok)
	}

	user := r.Form.Get("user")
	// 判断是否传入相应的user 值 , 根据user 的长度进行判断
	if len(user) > 0 {
		fmt.Fprintf(w, "hello %s \n", user)
	} else {
		fmt.Fprintf(w, "请你按照格式输入 localhost:port?user=XXX \n")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":20080", nil)
}
