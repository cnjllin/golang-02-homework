package main

import "fmt"
import "time"

func sum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	fmt.Printf("start of %v\n", s)
	c <- sum // send sum to c
	fmt.Printf("end of %v\n", s)
}

func main() {
	//s := []string{"hello", "golang", "c++", "world"}
	s := []int{1, 2, 3, 4, 5, 6}

	c1 := make(chan int, 1)
	//c2 := make(chan int)
	go sum(s[:len(s)/2], c1)
	//go sum(s[len(s)/2:], c2)
	time.Sleep(time.Second)
	x, y := <-c1, 0 // receive from c

	fmt.Println(x, y, x+y)
	time.Sleep(time.Second)
}
