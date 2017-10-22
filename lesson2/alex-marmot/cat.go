package main

// cat localfile
// cat URL

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func errorHandler()(code int, content string) {
	code, content = 1, ""
	return code, content
}

func getHTTPResult(url string)(code int, content string) {
	res, err := http.Get(url)
	if err != nil {
		return errorHandler()
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return errorHandler()
	}
	code = 0
	content = string(body)
	return code, content
}

func printFile(name string) (code int, content string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return errorHandler()
	}
	code = 0
	content = string(buf)
	return code, content
}

func execute(arg string, ch chan<- string) {
	status, content := printFile(arg)
	if status == 1 {
		status, content = getHTTPResult(arg)
		if status == 1 {
			ch <- "please type right url or path for localfile"
		}else{
			ch <- content
		}
	}else{
		ch <- content
	}
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
