package main
import(
	"fmt"
	"log"
	"net/rpc"
	"bufio"
	"os"
//	"net"
)

func main() {
	ips:=[]string{"localhost:12345","localhost:12346"}
	//clients:=make([]interface{},len(ips))
	clients:=[]rpc.Client{}
	for _,ip:=range ips {
		client, err := rpc.Dial("tcp", ip)
		clients = append(clients, *client)
		if err != nil {
			log.Fatal(err)
		}
	}
		in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		inchan := make(chan []byte, len(ips))
		for i := 0; i <= len(ips)-1; i++ {
			inchan <- line
			go goroutine_clientrecv(clients[i], inchan)
		}
	}
/*		for{
		for _,client:=range clients{
			line,_,err:=in.ReadLine()
			if err!=nil{
				log.Fatal(err)
			}
			var reply string

			err=client.Call("Listener.GetLine",line,&reply)
			if err!=nil{
				log.Fatal(err)
			}
			fmt.Println(reply)*/
//			var reply string
//			getresult(client,line,&reply)
		}



func goroutine_clientrecv(client rpc.Client,inchan chan []byte){
	for{
			cmd:=<-inchan
			var reply string

			err:=client.Call("Listener.GetLine",cmd,&reply)
			if err!=nil{
				log.Fatal(err)
			}
			fmt.Println(reply)
			//			var reply string
			//			getresult(client,line,&reply)
		}
		}

func getresult(client rpc.Client,command []byte,c *string){
	err:=client.Call("Listener.GetLine",command,c)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(*c)
}