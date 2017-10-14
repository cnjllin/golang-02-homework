package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request){

	r.ParseForm()

	user := r.Form.Get("user")
	if  len(user) > 0 {
		fmt.Fprintf(w,"Hello %s.", user)
	}else{
		fmt.Fprintf(w,"hehe %s.", "Stranger")
	}
}

func main(){
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
