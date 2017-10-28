package main

import "fmt"

func main() {
	a := 1
	b := 1
	sum := 1
	for {
		a, b = b, a+b
		if a >= 100 {
			break
		}
		sum += a
		fmt.Println(a)
	}
	fmt.Println(sum)
}
