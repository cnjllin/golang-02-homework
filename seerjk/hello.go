package main

import "fmt"

//fmt format
//全路径引用
//import "github.com/xxx/xxx"

func main() {
	fmt.Println("Hello, world!")
	//print("ddddd")
	//不要直接用print，初始版本debug使用
	fmt.Printf("%d abc %s", 1111, "world") //支持format string, ANSI标准
}

