package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func printFileContent(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(buf))
}

func printWebContent(weburl string) {
	resp, err := http.Get(weburl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s: %v\n", weburl, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: cat <filePath/fileName> or cat <URL,example:https://www.baidu.com>")
		os.Exit(1)
	}
	flag := os.Args[1:]
	for i := 0; i < len(flag); i++ {
		if strings.Contains(flag[i], "https://") || strings.Contains(flag[i], "http://") {
			printWebContent(flag[i])
		} else {
			printFileContent(flag[i])
		}
	}
}
