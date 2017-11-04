package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func printFile(name string) {
	// 读取name的数据并打印到标准输出
}

func printHTTP(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	// 读取resp.Body的内容，并打印到标准输出
	//resp.Body
}

func main() {
	var name string
	name = os.Args[1]
	log.Print(name)
	//name := os.Args[1]
	if strings.HasPrefix(name, "http://") {
		// 读取HTTP内容
		printFile(name)
	} else {
		// 以文件的方式打开
		printHTTP(name)
	}
}
