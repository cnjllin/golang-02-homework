package main

import "time"
import "fmt"

func main() {

	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 1000)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Millisecond * 1500)
		c1 <- "1.5"
	}()
	go func() {
		time.Sleep(time.Millisecond * 2000)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
