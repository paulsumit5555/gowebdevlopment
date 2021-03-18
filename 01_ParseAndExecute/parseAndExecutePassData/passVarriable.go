package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	
	tpl = template.Must(template.ParseFiles("passVarriable.gohtml"))

}

func main(){

	err := tpl.ExecuteTemplate(os.Stdout,"passVarriable.gohtml","hello how are you????")

	if err != nil{

		log.Fatalln(err)
	}
}