package main

import (
	"os"
	"fmt"
	"net"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main(){
	service := "0.0.0.0:1202"
	addr,err := net.ResolveTCPAddr("tcp",service)
	checkError(err)
	listener,err := net.ListenTCP("tcp",addr)
	checkError(err)
	for{
		conn,err := listener.Accept()
		if err!=nil{
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn){
	defer conn.Close()
	
	var buf [512]byte
	for {
		n,err := conn.Read(buf[0:])
		if err!=nil{
			conn.Close()
			return
		}
		s:=string(buf[0:n])

		if s[0:2]=="CD"{
			chDir(conn,s[3:])

		}else if s[0:3]=="DIR"{
			listDir(conn)
		}else if s[0:3]=="PWD"{
			pwd(conn)
		}
	}
}

func chDir(conn net.Conn,s string){
	if os.Chdir(s) == nil{
		conn.Write([]byte("OK"))
	}else{
		conn.Write([]byte("Error"))
	}
}

func pwd(conn net.Conn){
	s,err := os.Getwd()
	if err!=nil{
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

func listDir(conn net.Conn){
	defer conn.Write([]byte("\r\n"))
	dir,err := os.Open(".")
	if err!=nil{
		return
	}
	names,err := dir.Readdirnames(-1)
	for _,nm := range names{
		conn.Write([]byte(nm+"\r\n"))
	}

}

func checkError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal: %s\n",err.Error())
		os.Exit(1)
	}
}