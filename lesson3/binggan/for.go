package main

import "fmt"

func main() {
	s := "hello"
	for i, c := range s {
		fmt.Println(i, c)
	}

	for _, c := range s {
		fmt.Println(c)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var running = true
	for running {
		running = false
	}

	for {
		break
	}
}
