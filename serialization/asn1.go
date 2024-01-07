package main
import (
	"fmt"
	"os"
	"encoding/asn1"
	"time"
)

func checkError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal: %s\n",err.Error())
		os.Exit(1)
	}
}

func main(){
	mdata,err := asn1.Marshal(13)
	checkError(err)

	var n int 
	_,err1 := asn1.Unmarshal(mdata,&n)
	checkError(err1)
	fmt.Println("After Marshal/Unmarshal ",n)

	s := "hello"
	mdata, _ = asn1.Marshal(s)

	var newstr string
	asn1.Unmarshal(mdata, &newstr)
	fmt.Println("After Marshal/Unmarshal ",newstr)

	t := time.Now()
	mdata, err = asn1.Marshal(t)

	var newtime = new(time.Time)
	_, err1 = asn1.Unmarshal(mdata, &newtime)
	fmt.Println("After Marshal/Unmarshal ",newtime)
}