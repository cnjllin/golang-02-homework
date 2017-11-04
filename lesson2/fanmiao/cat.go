//实现一个cat命令，支持打印文件内容和http链接内容
//cat /home/work/xxx.txt
//cat http://www.baidu.com
//cat xxx.txt http://www.baidu.com
package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

//实现打印文件的功能
func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
//实现打印网页的功能
func printWeb(name string){
	resp, _ := http.Get(name)
	fmt.Println(resp)
}

func main() {
//抓取命令行参数
	for index,uri := range os.Args{
		if index == 0{
			continue
//判断命令行参数是文件还是网页，匹配http开头的为网页，默认为文件
		}else if strings.HasPrefix(uri,"http://"){
			printWeb(uri)
		}else{
                        printFile(uri)
		}
		fmt.Println("-------------------------------------")
	}
}
