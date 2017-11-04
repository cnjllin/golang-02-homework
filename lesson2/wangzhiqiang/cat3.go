package main

import (
	"os"
	"log"
	"io"
	"net/http"
	"fmt"
	"strings"
)

func GetString(a string){
	r,err := os.Open(a)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	io.Copy(os.Stdout,r)
}

func GetHttp(urls string){
	r,err := http.Get(urls)
	if err != nil{
		log.Fatal(err)
	}
	defer r.Body.Close()
	io.Copy(os.Stdout,r.Body)
}

func main(){
	if len(os.Args) ==1 {
		fmt.Println("请正确输入")
	}
	s := os.Args[1]
	if strings.HasPrefix(s,"http"){
		GetHttp(s)
	} else{
		GetString(s)
	}
}