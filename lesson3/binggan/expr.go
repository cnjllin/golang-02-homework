package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	stra := os.Args[1]
	op := os.Args[2] // "+", "-", "*", "/"
	strb := os.Args[3]
	// 解析字符串为数字
	// 根据操作符进行运算输出
	var a, b int
	var err error
	if a, err = strconv.Atoi(stra); err != nil {
		log.Fatal(err)
	}

	if b, err = strconv.Atoi(strb); err != nil {
		log.Fatal(err)
	}

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	default:
		fmt.Println("Invalid op " + op)
	}
}
