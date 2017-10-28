package main

import (
	"flag"
	"log"
	"os"
)

var (
	isPrintLine = flag.Bool("-l", false, "print lines")
)

func main() {
	flag.Parse()

	name := flag.Arg(0)
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // 关闭句柄，关闭socket, 关闭http的resp.Body, unlock锁
	//body, _ := ioutil.ReadAll(f)
	//fmt.Print(string(body))

}
