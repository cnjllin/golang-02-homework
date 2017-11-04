package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		// a int
		// b float32 // float64
		// c string
		// d bool
		// e byte
		// f rune

		// h int32  // int8, int16, int64
		// i uint32 // uint8, uint16,
		x int8
		y int
	)

	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(y))
}
