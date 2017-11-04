// 代码不简练, 逻辑不清楚, 没亮点
// 11-02 有问题
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	infos, err := ioutil.ReadDir("/proc")
	if err != nil {
		//log.Println(err)
	} else {
		for _, info := range infos {
			if info.IsDir() {
				f, err := os.Open("/proc/" + info.Name() + "/cmdline")
				if err != nil {
					//log.Println(err)
				} else {
					ff, err := ioutil.ReadAll(f)
					if err != nil {
						//log.Println(err)
					} else if string(ff) != "" {
						fmt.Printf("%s\n", info.Name()+"\t"+string(ff))
					}
				}
			}
		}
	}

}
