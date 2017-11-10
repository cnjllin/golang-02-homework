package main
import (
	"fmt"
//	"log"
	"net/http"
//	"encoding/json"
//	"os"
//	"io/ioutil"
	"strings"
//	"os"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
)

func Write2disk(s string,t string)string{
	body,err:=ioutil.ReadFile("homework.db")
	if err !=nil{
		log.Fatal(err)
	}
	f,err1:=os.Create("homework.db")
	if err1 !=nil{
		log.Fatal(err1)
	}
	defer f.Close()
	f.WriteString(string(body))
	maps:=make(map[string]string)
	maps[s]=t
	content,err1:=json.Marshal(maps)
	if err1 !=nil{
		log.Fatal(err)
	}
	f.WriteString(string(content))
	return "ok"
}

func Read2http(s string)string{
	f, err := os.Open("homework.db")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, err1 := ioutil.ReadAll(f)
	if err1 != nil {
		log.Fatal(err1)
	}
	content := []byte(body)
	maps := make(map[string]string)
	json.Unmarshal(content, &maps)
	key := s
	for i, _ := range maps {
		if i == key {
			//log.Println(maps[key])
			return maps[key]
		}
	}
	return "False"
}
func handlerx(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Fprintf(w,"%s","method get or set you can do ")
	//fmt.Fprintf(w,"%s",Read2http(key))

}
func handler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	str:=r.URL.Path //
	s:=strings.Split(str,"/")
	key:=s[2:][0]
	if len(s)>3{
		fmt.Fprintln(w,"your key error")
	}else if Read2http(key) =="False"{
		fmt.Fprintln(w,"can't found your key from data")    //
	} else {
		//fmt.Fprintf(w,"%s",s[2:][0])
		fmt.Fprintf(w,"%s",Read2http(key))
	}
}
func handler2(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	s:=strings.Split(r.URL.Path,"/")
	if len(s)!=4{
		fmt.Fprintln(w,"pls input standard url,like: 192.168.1.1:8000/set/key/value")
	}else if Read2http(s[2:3][0]) == "False"{
		//fmt.Fprintf(w,Read2http(s[2:3][0]))
		key:=s[2:3][0]
		value:=s[3:][0]
		result:=Write2disk(key,value)
		if result !="ok"{
			fmt.Fprintln(w,"save map faild,pls consult manager")
		}else{
			fmt.Fprintf(w,"SAVE key:%s,value:%s",key,value)
		}

	}else{
		fmt.Fprintf(w,"the key %s has exist",s[2:3][0])
	}
}
func main(){
	http.HandleFunc("/",handlerx)
	http.HandleFunc("/get/",handler)
	http.HandleFunc("/set/",handler2)
	http.ListenAndServe(":8000",nil)
}

