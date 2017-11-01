package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	x := 1000
	y := 100
	percent := float32(y) / float32(x) * 100
	fmt.Printf("used:%v%%\n", percent)

	s := "a12"
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("x=", x)
}
