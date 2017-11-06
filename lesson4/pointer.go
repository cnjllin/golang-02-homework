package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

func main() {
	s1 := Student{
		Id:   1,
		Name: "alice",
	}
	fmt.Println(s1)
	var p *Student
	p = &s1
	p.Id = 2
	fmt.Println(s1)
	var p1 *int
	p1 = &s1.Id
	*p1 = 3
	fmt.Println(s1)
}
