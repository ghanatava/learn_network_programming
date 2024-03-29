package main
import (
	"fmt"
	"os"
	"net"
)

func main(){
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"Usage: %s dotted-ip-addr\n")
		os.Exit(1)
	}
	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)

	if addr == nil{
		fmt.Println("Invalid Address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones,bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	os.Exit(0)
}