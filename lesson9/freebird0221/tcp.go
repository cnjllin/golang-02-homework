package main

import (
  "net"
  "fmt"
  "log"
  "os"
  "io"
  "strings"
  "bufio"
  // "io/ioutil"
  "regexp"
  "path/filepath"
)

type Request struct {
  Method string
  Path string
  Headers map[string]string
  Args map[string]string
}

var allow_ext = [...]string{".jpg", ".jpeg", ".png", ".webp"}

var BaseDir string
//MIME
var http_200 = `HTTP/1.1 200 OK
Content-Type: image/jpeg
Content-Length: %d
Connection: keep-alive
Server: Reboot

`

var http_index = `HTTP/1.1 200 OK
Content-Type: text/html; charset=utf-8
Content-Length: %d
Connection: keep-alive
Server: Reboot

<html>
<body>
<div>
<p>%02d images</p>
<p>
%s</p>
</div>
</body>
</html>
`


var http_404 = `HTTP/1.1 404 Not Found
Content-Type: image/jpeg
Content-Length: 0
Connection: keep-alive
Server: Reboot
`

func handleIdx(conn net.Conn, url Request) {
  imgs_fd, err := os.Open(filepath.Join(BaseDir, "imgs"))
  if err != nil {
    log.Fatal(err)
  }

  names, _ := imgs_fd.Readdirnames(-1)
  if err != nil {
    log.Fatal(err)
  }
  imgs := make([]string, 0, len(names))
  for _, name := range names {
    file_ext := filepath.Ext(name)
    for _, ext := range allow_ext {
      if ext == file_ext {
        // 构建 <a><img>
        imgs = append(imgs, fmt.Sprintf("<p><a href=\"/imgs/%s\"><img src=\"/imgs/%s\"></a></p>\n", name, name))
      }
    }
  }
  img_tags := strings.Join(imgs, "")
  fmt.Fprintf(conn, http_index, 69 + len(img_tags), len(names), img_tags)
}

func handleConn(conn net.Conn) {
  defer conn.Close()

  url := parse_url(conn)
  if url.Path == "/" {
    handleIdx(conn, url)
  } else {
    handleImg(conn, url)
  }
}

func handleImg(conn net.Conn, url Request){
  // get file name
  fields := strings.Split(url.Path, "/")
  filename := fields[len(fields) - 1]

  // open File or 404
  fd, err := os.Open(filepath.Join(BaseDir, "imgs", filename))
  if err != nil {
    log.Println(err)
    fmt.Fprintf(conn, http_404)
    return
  }
  defer fd.Close()
  info, err := fd.Stat()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Fprintf(conn, http_200, info.Size())
  io.Copy(conn, fd)
}

func parse_url(conn net.Conn) Request {
  url := Request {
    Method: "",
    Path: "",
    Args: make(map[string]string),
    Headers: make(map[string]string),
  }
  reader := bufio.NewReader(conn)
  for i := 0; true; i++ {
    byte_line, _, _ := reader.ReadLine()
    line := string(byte_line)
    //fmt.Println(line)

    // 首行 解析
    if i == 0 {
      fields := strings.Fields(line)
      url.Method = fields[0]
      path_args := strings.Split(fields[1], "?")
      url.Path = path_args[0]

      // 解析URL参数
      if len(path_args) > 1 {
        s := regexp.MustCompile("a*").Split(path_args[1], -1)
        for j := 0; true; j += 2 {
          url.Args[s[j]] = s[j + 1]
        }
      }
    } else {

      if string(line) == "" {
        break
      }
      k_v := strings.Split(line, ": ")
      url.Headers[k_v[0]] = k_v[1]
    }
  }
  return url
}

func main() {
  // set base dir
  ex, _ := os.Executable()
  BaseDir = filepath.Dir(ex)
  //BaseDir = "/Users/marcus/Projects/go/src" // 测试使用, go run BaseDir 位于/tmp 无法访问 imgs 文件夹

  add := ":8021"
  listener, err := net.Listen("tcp", add)
  fmt.Println("Listen" + add)
  if err != nil {
    log.Fatal(err)
  }
  defer listener.Close()

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatal(err)
    }
    go handleConn(conn)
  }
}