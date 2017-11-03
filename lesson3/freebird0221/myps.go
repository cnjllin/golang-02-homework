package main

import (
  "fmt"
  "log"
  "os"
  "strconv"
  "path/filepath"
  "io/ioutil"
)

const (
  BASE_PATH = "/proc"
)

func main(){
  fd, _ := os.Open(BASE_PATH)
  infos, _ := fd.Readdir(-1)
  for _, info := range infos {
    if _, err := strconv.Atoi(info.Name()); err != nil || !info.IsDir() {
      continue
    }
    filepath := filepath.Join(BASE_PATH, info.Name(), "cmdline")    
    fd, err := os.Open(filepath)
    hundle_err(err)
    cmdline, err := ioutil.ReadAll(fd)
    hundle_err(err)
    fmt.Printf("%s\t%s\n", info.Name(), cmdline)
  }
}

func hundle_err(err error) {
  if err != nil {
    log.Fatal(err)
  }
}