package main

import "fmt"

func main() {
	// 典型数组初始化赋值
	var a [3]int = [3]int{1, 2, 3}
	var ap = &a   // 取数组指针赋值给ap
	ap[1] = 8     // 对ap指向的数组进行修改
	(*ap)[2] = 88 // 完全等价于上一行

	// 初始化一个隐式声明长度的数组
	b := [...]int{1, 11, 111}
	// 初始化一个Slice(切片)
	c := []int{1, 11, 1111111}
	// 规定下标初始化切片
	d := []int{0: 1, 10: 11, 12: 111}
	d[1] = 123
	e := d //切片赋值,浅拷贝
	e[1] = 456
	var f [3]int
	f = a // 数组赋值,深拷贝
	f[1] = 888
	//b = c
	copy(b[:], c) // 把切片赋值给数组的方式
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", ap, ap)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	e[2] = 999
	fmt.Printf("%v, %p, %T\n", d, d, d)
	fmt.Printf("%v, %p, %T\n", e, e, e)
	fmt.Printf("%v, %p, %T\n", f, &f, f)
}
