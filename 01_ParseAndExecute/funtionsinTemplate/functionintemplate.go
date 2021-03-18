package main

import (
	"log"
	"os"
	"text/template"
	"strings"
)

var fm = template.FuncMap{
	"uc" : strings.ToUpper,
	"lc" : strings.ToLower,

}

var tpl *template.Template

func init(){
	
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("functionintemplate.gohtml"))

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
	
	err := tpl.ExecuteTemplate(os.Stdout,"functionintemplate.gohtml",p)

	if err != nil{

		log.Fatalln(err)
	}
}