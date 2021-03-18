package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	
	tpl = template.Must(template.ParseFiles("struct.gohtml"))

}

type person struct{
	Fname string
	Lname string
	Age int
}

func main(){

	p := person{
		Fname : "sumit",
		Lname : "paul",
		Age : 22,
	}
	
	err := tpl.ExecuteTemplate(os.Stdout,"struct.gohtml",p)

	if err != nil{

		log.Fatalln(err)
	}
}