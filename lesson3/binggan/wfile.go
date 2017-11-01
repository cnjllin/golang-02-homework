package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("hello\n")
	fmt.Fprintf(f, "%d x %d = %d\n", 3, 4, 3*4)
}
