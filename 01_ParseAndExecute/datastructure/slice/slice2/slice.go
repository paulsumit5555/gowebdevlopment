package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	
	tpl = template.Must(template.ParseFiles("slice.gohtml"))

}

func main(){

	slice := []string{"sumit","raj","rohan","kapil"}
	err := tpl.ExecuteTemplate(os.Stdout,"slice.gohtml",slice)

	if err != nil{

		log.Fatalln(err)
	}
}