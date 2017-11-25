package main

import "fmt"
import "time"
import "reflect"

func main() {
	s := "hello\n"
	for i, _ := range s {
		const k = 100
		time.Sleep(time.Duration(i+1) * time.Millisecond)
		time.Sleep(k * time.Millisecond)
		fmt.Println("c:", reflect.TypeOf((i + 1)))
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(reflect.TypeOf(100))
}
