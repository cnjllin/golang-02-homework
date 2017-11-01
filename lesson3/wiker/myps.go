package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	infos, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if info.IsDir() {
			pid, err := strconv.Atoi(info.Name())
			if err != nil {
				continue
			}
			path := "/proc/" + info.Name() + "/cmdline"
			inputfile, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(pid, string(inputfile))
		}
	}
}
