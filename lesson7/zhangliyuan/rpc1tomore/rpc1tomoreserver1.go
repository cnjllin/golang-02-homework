package main
import(
//	"fmt"
//	"fmt"
//	"os/exec"
	"log"
	"net"
	"net/rpc"

	"fmt"
	"os/exec"
)
type Listener int

func (l *Listener)GetLine(line []byte,ack *string)error{
	fmt.Println(string(line))
	cmd:=exec.Command("sh","-c",string(line))
	//cmd.Run()
	var out []byte
	var err error
	if out,err =cmd.CombinedOutput();err!=nil{
		log.Fatal(err)
	}
	result:=string(out)
	*ack=result
	return nil
}


func main() {
	addr,err:=net.ResolveTCPAddr("tcp","0.0.0.0:12345")
	if err!=nil{
		log.Fatal(err)
	}
	inbound,err:=net.ListenTCP("tcp",addr)
	if err!=nil{
		log.Fatal(err)
	}
	listeners:=new(Listener)
	rpc.Register(listeners)
	rpc.Accept(inbound)
}