package main

import(
	"net/http"
	"html/template"
	"log"
)


type handler int

var tpl *template.Template

 func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request){

 	err := r.ParseForm()
 	if err != nil {
 		log.Fatalln(err)
 	}


 	tpl.ExecuteTemplate(w,"request.gohtml",r.Form)
 }

 func home(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
 		if err != nil {
 			log.Fatalln(err)
 		}
	tpl.ExecuteTemplate(w,"request.gohtml",nil)
 }

 func hello(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
 		if err != nil {
 			log.Fatalln(err)
 		}
	tpl.ExecuteTemplate(w,"hello.gohtml",r.Form)
 }

func init(){
tpl = template.Must(template.ParseFiles("request.gohtml","hello.gohtml"))
}

func main(){
	http.HandleFunc("/home",home)
	http.HandleFunc("/hello",hello)
	http.ListenAndServe(":8080",nil)

}