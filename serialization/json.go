package main
import (
	"fmt"
	"os"
	"encoding/json"
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

func main(){
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
				Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}
	saveJson("person.json",person)
}

func saveJson(filename string,key interface{}){
	outFile,err := os.Create(filename)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal: %s\n",err.Error())
		os.Exit(1)
	}
}