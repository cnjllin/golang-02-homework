package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

func cat_http(i int) {
	r, err := http.Get(os.Args[i])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func cat_file(i int) {
	buf, err := ioutil.ReadFile(os.Args[i])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("请输入文件路径或网址")
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i][:4] == "http" {  //文件名需要大于4个字符
			cat_http(i)
		} else {
			cat_file(i)
		}
	}
}
