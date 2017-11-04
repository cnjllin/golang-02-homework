package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strconv"
)

func num_file(nums string)  {
	file := "/proc/" + nums  + "/cmdline"
	file_rep,file_err := ioutil.ReadFile(file)
	if file_err != nil {
		log.Fatal(file_err)
	}
	if string(file_rep) != "" {
		fmt.Printf("%s \t %s\n",nums,file_rep)
	}
}

func main()  {
	infos ,err :=ioutil.ReadDir("/proc")
		if err != nil{
			log.Fatal(err)
		}
	fmt.Printf("PID \t CMD\n")
		for _,info := range infos{
			if info.IsDir() {
				num,num_err := strconv.Atoi(info.Name())
				if num_err == nil{
					strnum:=strconv.Itoa(num)
					num_file(strnum)
				}

			}
		}
}

