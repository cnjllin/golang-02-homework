```go
package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	m, _ := url.ParseQuery(u.RawQuery)
	value, ok := m["user"]
	if ok {
		fmt.Fprintf(w, "hello %s.\n", value[0])
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8800", nil)
}
```