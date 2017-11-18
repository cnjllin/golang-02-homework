package main

import(
	"os"
	"fmt"
	//"io"
	"bufio"
	"log"
	"os/exec"
	"strings"
	"io/ioutil"
)

func main(){
	hostname,_:=os.Hostname()
	dir,_:=os.Getwd()
	//ostype:=os.
	notice:=fmt.Sprintf("%s#%s  ",dir,hostname)

	readline:=bufio.NewScanner(os.Stdin)
	for{
		r,w,_:= os.Pipe()
		fmt.Print(notice, " ")
		if !readline.Scan() {
			log.Fatal()
		}
		line := readline.Text() //
		if len(line) == 0 {
			continue
		}
		var filename string
		fmt.Print(filename)
		/////////输入搞不定啊aa
		if strings.Contains(line,"<") {
			args_in := strings.Split(line, "<")
			line = args_in[0]
			filename = args_in[1]
			fd, _ := ioutil.ReadFile(filename)
			args := strings.Fields(line)
			w.Write(fd)
			var tmp []byte
			r.Read(tmp)
			cmd := exec.Command(args[0], args[1:]...)  //输入不知道怎么添加到命令参数中
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

		} else if strings.Contains( line,">") {
			args_out := strings.Split(line, ">")
			line = args_out[0]
			filename = args_out[1]
			args := strings.Fields(line)
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = w
		//	cmd.Stderr = os.Stderr
			var tmp []byte
			r.Read(tmp)
			r.Close()
			w.Close()
			fd, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
			defer fd.Close()
			fd.WriteString(string(tmp))

		} else if strings.Contains( line,"|") {
			args_out := strings.Split(line, "|")
//			for i:=0;i<len(args_out);i++{
//				s:=fmt.Sprintf("command%s",strconv.Itoa(i))
//				s=args_out[i]
//				t:=fmt.Sprintf("cmd%s",strconv.Itoa(i))
//				t=exec.
//			}
			command0 := args_out[0]
			command1 := args_out[1]
			args_0 := strings.Fields(command0)
			args_1 := strings.Fields(command1)
			cmd_0 := exec.Command(args_0[0], args_0[1:]...)
			cmd_1 := exec.Command(args_1[0], args_1[1:]...)
			cmd_0.Stdin = os.Stdin
			cmd_0.Stdout = w
			cmd_1.Stdin = r
			cmd_1.Stdout = os.Stdout
			cmd_0.Start()
			cmd_1.Start()
			cmd_0.Wait()
			r.Close()
			w.Close()
			cmd_1.Wait()
			

		} else {
			args := strings.Fields(line)
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

