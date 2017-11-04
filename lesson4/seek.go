package main

import (
	"fmt"
	"os"
)

func main() {
	bytes := []byte{1, 1, 1, 1}
	f, _ := os.Create("1.db")
	defer f.Close()
	f.Seek(99999999999, 0) //相当于开始位置偏移1
	_, err := f.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}
