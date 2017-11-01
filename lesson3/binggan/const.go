package main

import "fmt"

const PI = 3.14
const StatusOk = "OK"

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
)

const (
	RED = iota
	BLUE
	GREEN
)

func main() {
	fmt.Println(PI)
	fmt.Println(RED)
	fmt.Println(BLUE)
	fmt.Println(GREEN)
}
