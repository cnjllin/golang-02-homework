package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const (
	DIR = "/proc/"
)

// 获取/proc/目录下所有的pid目录名
func getPidDirs() []string {
	reg, _ := regexp.Compile(`(\d+)`)
	infos, err := ioutil.ReadDir(DIR)

	if err != nil {
		log.Fatal(err)
	}

	var pidDirs []string
	for _, info := range infos {
		if info.IsDir() && reg.MatchString(info.Name()) {
			pidDirs = append(pidDirs, info.Name())
		}
	}
	//sort.Strings(pidDirs)
	return pidDirs
}

// 按pid目录名升序排序
func sortArrIntstr(tarstr []string) []string {
	var arrint []int
	for _, str := range tarstr {
		a, _ := strconv.Atoi(str)
		arrint = append(arrint, a)
	}
	sort.Ints(arrint)

	var resstr []string
	for _, i := range arrint {
		s := strconv.Itoa(i)
		resstr = append(resstr, s)
	}
	return resstr
}

// 获取pid目录下cmdline文件内容
func getCmmand(filepath string) string {
	filepath += "cmdline"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	body, _ := ioutil.ReadAll(f)

	return string(body)
}

func main() {
	for _, pidDir := range sortArrIntstr(getPidDirs()) {
		filepath := DIR + pidDir + "/"
		cmd := getCmmand(filepath)

		// fmt.Printf("%s\t\t%.60s\n", pidDir, cmd)
		if cmd != "" { //过滤cmdline为空的
			fmt.Printf("%-10s%.60s\n", pidDir, cmd)
		}

	}
}
