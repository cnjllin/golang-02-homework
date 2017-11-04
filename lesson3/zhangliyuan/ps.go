package main
import (
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	"log"
)

func Readfile(pid int)string{
	f,err:=os.Open(fmt.Sprintf("/proc/%d/task/%d/comm",pid,pid))
	if err!=nil{
		log.Fatal(err)
	}
	body,_:=ioutil.ReadAll(f)
	//fmt.Println(string(body))
	defer f.Close()
	return string(body)
}
//func Write2file(pid int,content string){
//	
//}
func main(){
	f,err:=ioutil.ReadDir("/proc")
	if err != nil{
		log.Fatal(err)
		}
	for _,c:=range f{
		s,err:=strconv.Atoi(c.Name())
		//fmt.Println(c.Name())
		if err !=nil{
			continue
		}else{
		//	fmt.Println(s)
			content:=Readfile(s)
		//	fmt.Println(content)
		//	fmt.Printf("%s\t%s",string(s),content)
			fmt.Printf("%d\t%s",s,content)
			}
		}
//			file,_:=os.Create("a.txt")
//			defer file.Close()
//			file.WriteString(s)
}

