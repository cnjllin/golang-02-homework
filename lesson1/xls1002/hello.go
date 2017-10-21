package main

import (
	"fmt"
	"log"
	"net/http"
)

// http://localhost:8800/?user=pc
// 响应一个 hello pc

func parseParam(w http.ResponseWriter, r *http.Request) {

	/* 解析参数 */
	err := r.ParseForm()
	if err != nil {
		log.Fatal("参数解析错误:", err)
	}

	/* 获取参数中的user值 */
	userValue := r.Form.Get("user")
	if len(userValue) == 0 {
		fmt.Fprintf(w, "<h1>请输入user参数的值!</h1>")
	} else {
		fmt.Fprintf(w, "<h1>hello %s</h1>", userValue)
	}
}

func main() {

	/* 设置路由 */
	http.HandleFunc("/", parseParam)

	/* 监听8800端口 */
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		log.Fatal("监听8800端口失败：", err)
	}
}
