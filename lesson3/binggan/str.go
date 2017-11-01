package main

import "fmt"

type Student struct {
	Name string `json:"name"`
}

func reverse(s string) string {
	s1 := []rune(s)
	// 翻转数组s1
	s2 := make([]rune, len(s1))
	for i := 0; i < len(s1); i++ {
		s2[i] = s1[len(s1)-i-1]
	}
	return string(s2)
}

func main() {
	doc := `
hello world
即使换行也不影响\n
`
	fmt.Print(doc)

	s := "hello"

	s = "a" + s[1:]
	fmt.Println(s)

	s1 := []byte(s)
	s1[0] = 'a'
	s = string(s1)

	s = "你好，世界"
	s2 := []rune(s)
	s = string(s2)
	fmt.Println(s)

	x, y := 1, 2
	x, y = y, x
}
