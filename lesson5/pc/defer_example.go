package main

import (
	"log"
)

/*
2017/11/11 16:55:52 print
2017/11/11 16:55:52 defer2
2017/11/11 16:55:52 defer1
2017/11/11 16:55:52 main
*/

func print() {
	defer log.Println("defer1")
	defer log.Println("defer2")
	log.Println("print")
}

func main() {
	print()
	log.Println(1 << 2)
	log.Println(10 << 2)
}
