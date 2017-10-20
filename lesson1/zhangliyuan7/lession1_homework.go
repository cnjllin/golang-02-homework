package main

import (
	"net/http"
	"fmt"
//	"log"
	"time"
//	"net"

)

func handler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Fprintf(w,"Time now : %s",time.Now().String())
//	fmt.Fprintf(w,"Hello %s",html.EscapeString(r.URL.Path[1:]))
	if r.Method=="GET" && r.Form["username"] !=nil {
		fmt.Fprintf(w,"Hello %s",r.Form["username"][0])
	}
}
func main(){
/*	fmt.Println("hello world")
	request url:"http://localhost:8800/?username=pc"
 */

 	http.HandleFunc("/",handler)
 	http.ListenAndServe(":8080",nil)


}
