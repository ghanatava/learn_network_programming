package main

import (
	"fmt"
	"os"
	"bytes"
	"encoding/asn1"
	"io"
	"net"
	"time"
)

func main(){
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"Usage: %s host:port\n",os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	addr,err := net.ResolveTCPAddr("tcp4",service)
	checkError(err)
	conn,err := net.DialTCP("tcp",nil,addr)
	checkError(err)

	result,err := readFully(conn)
	checkError(err)

	var newtime time.Time
	_,err1 := asn1.Unmarshal(result,&newtime)
	checkError(err1)

	fmt.Println("After Marshal/Unmarshal: ",newtime.String())
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
			os.Exit(1)
	}
}

func readFully(conn net.Conn)([]byte,error){
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n,err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err!=nil{
			if err == io.EOF{
				break
			}
			return nil,err
		}
	}
	return result.Bytes(), nil
}