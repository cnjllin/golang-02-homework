package main

import (
  "fmt"
  "os"
  "io"
  "bufio"
  "strings"
  "os/exec"
  "os/user"
)

type MyCmd struct {
  cmd *exec.Cmd
  pipe_read  *io.PipeReader
  pipe_write *io.PipeWriter
}

func main() {
  var user_name string
  user, _ := user.Current()
  user_name = user.Name
  // fmt.Println("User Name:")
  // fmt.Scanln(&user_name)
  host, _ := os.Hostname()
  prompt := fmt.Sprintf("[%s@%s]", user_name, host)
  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print(prompt)
    if !scanner.Scan() {
      break
    }
    line := scanner.Text()
    if line == "exit" {
      return
    }
    cmd_strs := strings.Split(line, "|")
    length := len(cmd_strs)
    cmds := make([]MyCmd, length)
    for i, cmd_str := range cmd_strs {
      cmd_args := strings.Fields(strings.TrimSpace(cmd_str))
      cmd := exec.Command(cmd_args[0], cmd_args[1:]...)
      pipe_read, pipe_write := io.Pipe()

      if i == 0 {
        cmd.Stdin = os.Stdin
      } else {
        cmd.Stdin = cmds[i - 1].pipe_read
      }
      if i == length - 1 {
        cmd.Stdout = os.Stdout
      } else {
        cmd.Stdout = pipe_write
      }
      my_cmd := MyCmd{
        cmd: cmd,
        pipe_read: pipe_read,
        pipe_write: pipe_write,
      }
      cmds[i] = my_cmd
      my_cmd.cmd.Start()
    }
    for _, my_cmd := range cmds {
      my_cmd.cmd.Wait()
      my_cmd.pipe_read.Close()
      my_cmd.pipe_write.Close()
    }
  }
}
