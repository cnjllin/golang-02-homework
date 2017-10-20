package main

import (
	"fmt"
)

/*
func main() {
	fmt.Println("hello word!")
}
*/
//常里赋值
// 关键字 常量名 = 常量值
// const  name = xxxx
/*
const (
	a = 1
	PI = 3.14
)
const a =1
const PI = 3.14
*/

//变量赋值
// 关键字  变量名  变量类型 = 变量值
// var  name type = xxxx
/*
var (
	a = 1
	b = 2
)
var a int 1
var b int 2
*/

//if语句示例
/*
func main() {
	a := 1
	if (10 / a) > 1 {
		fmt.Println("ok")
	}

}
*/

//for语句示例
/*
func main() {
	a := 1
	for {
		a++
		if a > 7 {
			break
		}
		fmt.Println(a)
	}
	fmt.Println("Over")
}



func main() {
	a := 1
	for a <= 3 {
		a++
		fmt.Println(a)
	}
}

func main() {
	a := 1
	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a)
	}
	fmt.Println("Over")
}
*/

//switch示例
/*
func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")

	case 1:
		fmt.Println("a=1")

	default:
		fmt.Println("None")
	}
}
*/

/*
func main() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
		fallthrough
	default:
		fmt.Println("None")
	}
}
*/

/*
func main() {
	switch a := 1; {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
}
*/

//跳转语句
/*
func main() {
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("OK")
}

func main() {
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL1
			}
		}
	}
LABEL1:
	fmt.Println("OK")
}
*/

/*
func main() {
LABEL1:
	for i := 0; i < 10; i++ {
		for {
			continue LABEL1
			fmt.Println("i")
		}
	}
	fmt.Println("OK")
}
*/
func main() {
	var a [2]int
	var b [1]int
	b = a
	fmt.Println(b)
}

