package main

import "fmt"

var (
	m map[string]string
)

func print() {
	var x int
	fmt.Println(x)
}

// 逃逸
// 堆和栈
func print1() *int {
	var x int
	return &x
}

func main() {
	print()

	p := print1()
	fmt.Println(*p)
	var x [10000000]int
}
