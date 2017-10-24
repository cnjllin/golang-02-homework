package main

import (
	"fmt"
	"io/ioutil"
	//	"log"
	"net/http"
	"os"
)

// 获取文件内容
func Get_file_message(str string) {
	buf, err := ioutil.ReadFile(str)
	// ReadFile 函数里面已经判断了是否为文件或目录路径
	if err != nil {
		// log.Println(err)
		return
	}
	fmt.Printf("%s", string(buf))
}

// 获取http 网页的内容
func Get_http_message(str string) {
	resp, err := http.Get(str)
	if err != nil {
		//log.Println(err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Println(err)
		return
	}
	fmt.Println(string(buf))

}

// 根据访问url 的http 状态码进行判断
func Get_http_status(str string) int {
	resp, err := http.Get(str)
	if err != nil {
		//log.Println(err)
		return 0
	}
	defer resp.Body.Close()
	//fmt.Println(resp.StatusCode)
	return resp.StatusCode

}
func main() {
	/*
		通过 url 的 http 状态码判断是否为url 或者是文件
		1. 存在的问题 如果url 输入错误，则会认为是文件
		通过是否为文件进行判断，则会出现 只输入目录路径的情况

		需要老师帮忙提一下 相应的思路
	*/

	s := os.Args[1:]

	for _, n := range s {
		code_status := Get_http_status(n)
		if code_status == 0 {
			Get_file_message(n)
		} else {
			Get_http_message(n)
		}
	}

}
