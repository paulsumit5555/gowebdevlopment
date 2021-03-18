package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	
	tpl = template.Must(template.ParseFiles("map.gohtml"))

}

func main(){

	m := map[string]string{
		"fname": "sumit",
		"lname": "paul",
	}
	err := tpl.ExecuteTemplate(os.Stdout,"map.gohtml",m)

	if err != nil{

		log.Fatalln(err)
	}
}