package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()
	for {
		//d2 := []byte{115, 111, 109, 101, 10}
		//n2, err := f.Write(d2)
		//check(err)
		//fmt.Printf("wrote %d bytes\n", n2)
		//	n3, err := f.WriteString("writes\n")
		//	fmt.Printf("wrote %d bytes\n", n3)
		//	//f.Sync()
		w := bufio.NewWriter(f)
		n4, _ := w.WriteString("bufferedfdasfdsafdsafdsafdsafdsafdsafdsfasdfsafdsafdsafdsafdsafdsa\n")
		fmt.Printf("wrote %d bytes\n", n4)
		time.Sleep(400 * time.Millisecond)
		w.Flush()
	}
}
