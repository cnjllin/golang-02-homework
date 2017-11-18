package main

import (
	"fmt"
	"math"
)

type Areaer interface {
	Area() float64
}

type 矩形 struct {
	l, w float64
}

type Circle struct {
	r float64
}

func (s 矩形) Area() float64 {
	return s.l * s.w
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	round := Circle{5.2}
	很正的矩形 := 矩形{3, 4}
	形状s := []Areaer{round, 很正的矩形}
	fmt.Println(形状s[0].Area()) // 求面积
	fmt.Println(形状s[1].Area()) // 求面积
}
