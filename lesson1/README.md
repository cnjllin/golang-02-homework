# lesson1 homework

```
package main

import (
	"fmt"
	"net/http"
)

// http://localhost:8800/?user=pc
// 响应一个 hello pc

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello %s!</h1>", r.URL)
}

func user_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>user %s!</h1>", r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user/", user_handler)
	http.ListenAndServe(":8800", nil)
}
```
