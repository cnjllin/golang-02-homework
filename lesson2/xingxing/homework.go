package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func paurl(url string) {
	resp, err := http.Get(url) //调试
	if err != nil {            // 调试
		fmt.Println(err) //调试
	} // 调试
	fmt.Println(resp)
}
func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
func main() {
	a := os.Args[1]
	b := os.Args[2]
	printFile(a)
	paurl(b)

}
