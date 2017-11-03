package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func printCmd(s string) string {
	dir := "/proc/" + s + "/cmdline"
	f, err := ioutil.ReadFile(dir)
	if err != nil {
		log.Fatal(err)
	}
	return string(f)
}

func findProc(s string) int {
	proc, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return proc
}

func main() {
	f, err := ioutil.ReadDir("proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, dirs := range f {
		fmt.Println(findProc(dirs.Name()), printCmd(dirs.Name()))
	}
}
