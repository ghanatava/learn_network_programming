package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Person struct {
	Name Name
	Email []Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct{
	Kind    string
	Address string
}

func (p Person)String()string{
	s := p.Name.Personal + p.Name.Family
	for _,v := range p.Email{
		s+="\n"+v.Kind+":"+v.Address
	}
	return s
}

func loadJson(filename string,key interface{}){
	inFile,err := os.Open(filename)
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func main(){
	var person Person
	loadJson("person.json",&person)
	fmt.Println("Person",person.String())
}

func checkError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal: %s\n",err.Error())
		os.Exit(1)
	}
}