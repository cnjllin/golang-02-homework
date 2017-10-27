package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func cat_http(i int) {
	r, err := http.Get(os.Args[i])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func cat_file(i int) {
	f, err := os.OpenFile(os.Args[i], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	f.Close()

}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("请输入文件路径或网址")
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i][:4] == "http" {
			cat_http(i)
		} else {
			cat_file(i)
		}
	}
}
