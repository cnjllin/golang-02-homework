
package main

import (
	"fmt"
	"io"
	"strings"
)

//!+bytecounter

type BCounter int

func (c *BCounter) Write(p []byte) (int, error) {
	fmt.Printf("got %v\n", string(p))

	return len(p), io.ErrShortWrite
}

//!-bytecounter

func main() {
	//!+main
	var c BCounter
	io.CopyN(&c, strings.NewReader("abcdefghijklmn"), 2)
}