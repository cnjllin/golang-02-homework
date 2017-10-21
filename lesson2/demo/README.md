# 作业

实现一个cat命令，支持打印文件内容和http链接内容

```bash
cat /home/work/xxx.txt
cat http://www.baidu.com
cat xxx.txt http://www.baidu.com
```


```golang
package main

import "net/http"
import "fmt"


func main() {
	resp, _ := http.Get("http://www.baidu.com")
	fmt.Println(resp)
}
```
