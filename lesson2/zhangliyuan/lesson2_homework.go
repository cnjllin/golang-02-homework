package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"net/http"
	"strings"
)

func readfile(filename string ){
	buf,err:=ioutil.ReadFile(filename)
	if err !=nil{
		log.Fatal()
		return
	}
	fmt.Println(filename)
	fmt.Println(string(buf))
	fmt.Println("#############################################################")
}

func readurl(urls string ){
	resp,err := http.Get(urls)
	if err != nil{
		log.Fatal(err)
		return
	}
	fmt.Println(resp)
	fmt.Println("#############################################################")
}

func main(){
	li:=os.Args[1:]
	for i:=0;i<len(li);i++{
		if strings.Contains(li[i],"https://") || strings.Contains(li[i],"http://"){
			readurl(li[i])
		} else {
			readfile(li[i])
		}
	}
}
