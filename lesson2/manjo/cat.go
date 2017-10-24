package main

//import "fmt"

//import "io/ioutil"
import "io"
import "os"
import "net/http"
import "log"
import "strings"

func catfile(s string) {
        src, err := os.Open(s)
        if err != nil {
                log.Print(err)
        }
        io.Copy(os.Stdout, src)
}
func cathttp(s string) {
        resp, err := http.Get(s)
        if err != nil {
                log.Print(err)
        }
        io.Copy(os.Stdout, resp.Body)
        //defer resp.Body.Close()
        //body, err := ioutil.ReadAll(resp.Body)
        //fmt.Printf("%v", string(body))
}
func main() {
        for _, file := range os.Args[1:] {
                if strings.HasPrefix(file, "http://") {
                        cathttp(file)
                } else {
                        catfile(file)
                }
        }

}
