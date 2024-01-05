package main

import (
	"fmt"
	"os"
	"net"
)

func main(){
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"Usage: %s hostname\n",os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr,err := net.ResolveIPAddr("ip",name)

	if err!=nil{
		fmt.Println("Resolution error")
		os.Exit(1)
	}
	fmt.Println("Ressolved Address: ",addr.String())
	os.Exit(0)
}