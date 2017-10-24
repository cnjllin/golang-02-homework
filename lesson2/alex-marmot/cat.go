package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getHTTPResult(url string)(string) {
	res, err := http.Get(url)
	if err != nil {
		return err.Error()
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func printFile(name string)(string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func execute(arg string, ch chan<- string) {
	var content string
	if strings.HasPrefix(arg, "http") {
		content = getHTTPResult(arg)
	} else {
		content = printFile(arg)
	}
	ch <- content
}

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage: ./cat args")
		return
	}

	ch := make(chan string)
	for _, arg := range os.Args[1:] {
      go execute(arg, ch)
  }

  for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}
