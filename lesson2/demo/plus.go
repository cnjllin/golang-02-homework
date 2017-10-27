package main

import "fmt"
import fl "flag"

// ./echo fsafdas
// fsafdas\n
// ./echo -n fsafdas
// fsafdas
var a = "aaaa"

func main() {
	linef := fl.Bool("n", false, "")
	fl.Parse()
	fmt.Println("linef", *linef)
	fmt.Println("args", fl.Args())
	fmt.Println("a", a)
}
