package main

// cat localfile
// cat URL

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getHTTPResult(url string)(code int) {
	res, err := http.Get(url)
	if err != nil {
		code = 1
		return code
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		code = 1
		return code
	}
	code = 0
	fmt.Println(string(body))
	return code
}

func printFile(name string) (code int) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		code = 1
		return code
	}
	code = 0
	fmt.Println(string(buf))
	return code
}

func execute(arg string) {
	result := printFile(arg)
	if result == 1 {
		result = getHTTPResult(arg)
	}
	if result == 1 {
		fmt.Println("please type right url or path for localfile")
	}
}

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage: ./cat args")
		return
	}

	for _, arg := range os.Args[1:] {
      execute(arg)
  }
}
