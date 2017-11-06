package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

type Teacher struct {
	Id   int
	Name string
	Sex  string
}

func main() {
	var s Student
	s.Id = 1
	s.Name = "jack"
	fmt.Println(s)
	s1 := Teacher{
		Id:   2,
		Name: "alice",
		Sex:  "male",
	}
	fmt.Println(s1)
	s1 = Teacher(s)
	fmt.Println(s1)
}
