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
	
	tpl = template.Must(template.ParseGlob("*.gohtml"))

}

type person struct{
	Fname string
	Lname string
	Age int
}

func main(){

	
	
	err := tpl.ExecuteTemplate(os.Stdout,"functionintemplate.gohtml",nil)

	if err != nil{

		log.Fatalln(err)
	}
}