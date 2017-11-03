package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"io/ioutil"
)


func main() {
	psMap := make(map[int]string)
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
	if pid, err := strconv.Atoi(file.Name()); err != nil {
			continue
		} else {
			dat, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline",pid))
			if err != nil || len(dat) == 0 {
				continue
			}
			cmdline := strings.Replace(string(dat), " ", "", -1)
			psMap[pid] = cmdline
		}
	}
	fmt.Println("  pid      cmdline")
	for p, c := range psMap {
		fmt.Printf("%5d      %-20s\n", p, c)	
	}

}
