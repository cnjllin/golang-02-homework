package main

import (
	"time"
	"fmt"
	"sync"
)

type IncType struct {
	i int64
	ILock sync.Mutex
}

var Incvar IncType

func Inc(inc *IncType) {
	counter := 0
	for {
		counter += 1
		inc.ILock.Lock()
		inc.i += 1
		//inc.ILock.Unlock()
		if counter >= 10000000 {
			break
		}
	}
	fmt.Println("finish")
	return
}

func main() {

	go Inc(&Incvar)
	go Inc(&Incvar)
	go Inc(&Incvar)

	time.Sleep(time.Second * 5)
	fmt.Println(Incvar.i)
	return
}
