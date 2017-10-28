package main

import "fmt"

var (
	x = 0
)

func print() {
	x := 10
	_ = x
}

func main() {
	// 1. 首先寻找同级作用域内是否有x
	// 2. 如果找不到往上一级作用域寻找x
	// 作用域从大到小 全局 > 函数 > if/for语句块
	x := 10
	if true {
		x := 100
		_ = x
		fmt.Println(x)
	}
	fmt.Println(x)
}
