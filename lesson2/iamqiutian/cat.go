package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		log.Fatal("User input error!")
		return
	}
	if len(args) == 2 {
		if strings.HasPrefix(args[1], "http://") {
			openUrl(args[1])
			return
		}
		openFile(args[1])
		return
	}
	if len(args) > 2 {
		for _, path := range args[1:] {
			if strings.HasPrefix(path, "http://") {
				openUrl(path)
			} else {
				openFile(path)
			}
		}
	}
}

func openFile(filepath string) {
	out, err := os.Open(filepath)
	errorCheck(err)
	_, err = io.Copy(os.Stdout, out)
	errorCheck(err)
}

func openUrl(uri string) {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(uri)
	errorCheck(err)
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	errorCheck(err)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
