package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func Iscmdline(pid string) {
	f, err := os.Open("/proc/" + pid + "/cmdline")
	checkErr(err)
	ff, err := ioutil.ReadAll(f)
	checkErr(err)
	if string(ff) != "" {
		fmt.Printf("%s\n", pid+"\t"+string(ff))
	} else {
		Iscomm(pid)
	}
}

func Iscomm(pid string) {
	f, err := os.Open("/proc/" + pid + "/comm")
	checkErr(err)
	ff, err := ioutil.ReadAll(f)
	checkErr(err)
	if string(ff) != "" {
		fmt.Printf("%s", pid+"\t"+string(ff))
	}
}

func checkErr(err error) {
	if err != nil {
		log.Printf("%s", err)
	}
}

func main() {
	infos, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if info.IsDir() {
			s := info.Name()
			true, err := regexp.Match("[0-9]+", []byte(s))
			checkErr(err)
			if true {
				Iscmdline(s)
			}
			continue
		}
	}
}
