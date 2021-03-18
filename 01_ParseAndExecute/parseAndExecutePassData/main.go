package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	
	tpl = template.Must(template.ParseFiles("first.gohtml"))

}

func main(){

	err := tpl.ExecuteTemplate(os.Stdout,"first.gohtml",43)

	if err != nil{

		log.Fatalln(err)
	}
}