package main

//作业
//
//实现一个cat命令，支持打印文件内容和http链接内容，支持以下方式
//cat /home/work/xxx.txt
//cat http://www.baidu.com
//cat xxx.txt http://www.baidu.com

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func catFile(filePath string) (result string, err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func catURL(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func printCatUsage() {
	content := "cat support print file content of URL content.\n"
	content += "cat usage:\n\n"
	content += "1. cat filePath\n"
	content += "2. cat URL\n"
	content += "3. cat filePath URL\n\n"
	content += "Examples:\n\n"
	content += "cat /home/work/xxx.txt\n"
	content += "cat http://www.baidu.com\n"
	content += "cat xxx.txt http://www.baidu.com\n"

	fmt.Println(content)
}

func main() {
	help := flag.Bool("h", false, "print cat usage")
	flag.Parse()

	if *help {
		printCatUsage()
		return
	}

	var output string

	for i := 0; i < len(flag.Args()); i++ {
		var tmp string
		var err error

		if strings.Contains(flag.Arg(i), "http://") {
			tmp, err = catURL(flag.Arg(i))
		} else {
			tmp, err = catFile(flag.Arg(i))
		}

		if err != nil {
			log.Fatal(err)
		}
		output += tmp
	}
	fmt.Println(output)

	//url_input := "http://www.baabcidu.com"
	//content, err := catURL(url_input)
	//if err != nil {
	//	// handle error
	//	log.Fatal(err)
	//}
	//fmt.Println(content)
	//
	//fmt.Println(strings.Contains(url_input, "http://"))

	//content, err := catFile("/tmp/if1.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(content)
}
