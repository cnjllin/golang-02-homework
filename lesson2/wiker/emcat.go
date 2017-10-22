package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
	"strings"
)

func main() {
	slice := os.Args[1:]
	for _, value := range slice {
		if strings.HasPrefix(value, "http") {
			fmt.Println(CatHttp(value))
		} else {
			fmt.Println(CatFile(value))
		}
	}
}

func CatFile(f string) string {
	input, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func CatHttp(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
