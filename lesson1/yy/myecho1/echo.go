package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("m", "333", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
}
