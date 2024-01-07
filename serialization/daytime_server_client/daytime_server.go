package main
import (
	"fmt"
	"os"
	"net"
	"time"
	"encoding/asn1"
)

func main(){
	service := ":1200"
	addr,err := net.ResolveTCPAddr("tcp",service)
	checkError(err)
	listener,err := net.ListenTCP("tcp",addr)
	checkError(err)
	for {
		conn,err:=listener.Accept()
		if err!=nil{
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn){
	defer conn.Close()

	daytime := time.Now()
	mdata,err := asn1.Marshal(daytime)
	checkError(err)
	conn.Write(mdata)
}

func checkError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal: %s\n",err.Error())
		os.Exit(1)
	}
}