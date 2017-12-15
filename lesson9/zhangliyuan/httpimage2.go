package httpimage
import(
	"net"
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
//	"net/http"
)
var Content=`HTTP/1.1 200 OK
Data:Sat,29 July 2017 6:30:42
Content-Type: text/html
Centent-Length: %d
Connection: Keep-Alive
Server:image.server


`
var Content2=`HTTP/1.1 200 OK
Data:Sat,29 July 2017 6:30:42
Content-Type: image/jpeg
Centent-Length: %d
Connection: Keep-Alive
Server:image.server


`
var str1=`<html>
<body>
`
var str3=`</body>
</html>
`
func Homepage(conn net.Conn,filelist []string){

	str2:=`<image src=%v>%s</image>`

	var div string
	for _,file:=range filelist{
		div+=fmt.Sprintf(str2,file,strings.Split(file,"/")[len(file)-1])
	}
	conn.Write([]byte(fmt.Sprintf(Content,len([]byte(str1+div+str3)))+str1+div+str3))
	conn.Close()
	}

func GetImage(conn net.Conn,file string){
	f,err:=os.Open(file)
	defer f.Close()
	if err!=nil{
		log.Fatal(err)
	}
	contentfile:=[]byte{}
	buf:=make([]byte,1024)
	for{
		n,_:=f.Read(buf)
		contentfile=append(contentfile,buf[:n]...)
	}
	conn.Write([]byte(fmt.Sprintf(Content2,len(contentfile))+str1+string(contentfile)+str3))
	conn.Close()
}

func main() {
	lisenner,err:=net.Listen("tcp",":9000")
	defer lisenner.Close()
	if err!=nil{
		log.Fatal(err)
	}
	for{
		conn,err:=lisenner.Accept()
		if err!=nil{log.Fatal(err)}
		filelist:=[]string{"/image/a.jpeg","/image/b.jpeg"}
		//maps:=Getimage(conn)
		//fmt.Println(conn.Read)
		buf:=bufio.NewReader(conn)
		getline,_,_:=buf.ReadLine()
		sliceget:=strings.Fields(string(getline))
		if len(sliceget)>2{
			GetImage(conn,sliceget[1])
		}
		go Homepage(conn,filelist)
	}

}