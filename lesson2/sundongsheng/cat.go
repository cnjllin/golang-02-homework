package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	//"fmt"
)

func catWeb(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return err.Error()
	}

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	//fmt.Printf("%s", buff)
	return string(buff)

}

func catFile(filename string) string {
	buff, err := ioutil.ReadFile(filename)

	if err != nil {
		return err.Error()
	}

	return string(buff)

}
func main() {
	cmd := os.Args
	cmdlen := len(os.Args)

	if cmdlen < 2 {
		println("input error!")
		return
	}

	for i := 1; i < cmdlen; i++ {
		if strings.HasPrefix(cmd[i], "http") {
			println(catWeb(cmd[i]))
		} else {
			println(catFile(cmd[i]))
		}
	}
}
