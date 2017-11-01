package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

func main()  {
	if len(os.Args) < 2 {
		fmt.Println("Usage:curl_file file or url")
	}else {
		for i := 1; i < len(os.Args); i++ {
			buf, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				reps, reps_err := http.Get(os.Args[i])
				if reps_err != nil {
					fmt.Printf("第 %d 个参数 %s is not file or url \n", i, os.Args[i])
				} else {
					fmt.Printf("第 %d 个参数为URL，内容为：",i)
					fmt.Println(reps)
					fmt.Printf("\n")
				}
			} else {
				fmt.Printf("第 %d 个参数为文件，内容为：",i)
				fmt.Println(string(buf))
				fmt.Printf("\n")
			}
		}
	}
}
