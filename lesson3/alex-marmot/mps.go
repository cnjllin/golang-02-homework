package main

import (
  "io/ioutil"
  "fmt"
  "strings"
  "regexp"
)

func errorHandler(err error)  {
  if err != nil {
    panic(err.Error())
  }
}

func getCmdline(pid string) string {
  path := strings.Join([]string{"/proc/", pid, "/cmdline"}, "")
  res, err := ioutil.ReadFile(path)
  errorHandler(err)
  return string(res)
}

func format(pid string) (string, string) {
  cmd := getCmdline(pid)
  return cmd, pid
}

func main()  {
  res, err := ioutil.ReadDir("/proc")
  errorHandler(err)

  for _, filename := range res {
    match, _ := regexp.MatchString("([a-z]+)", filename.Name())
    if !match {
      cmd, pid := format(filename.Name())
      fmt.Printf("%s \t %s \n",pid, cmd)
    }
  }
}
