package main

import (
	"github.com/julienschmidt/httprouter"
	//"gopkg.in/mgo.v2"
	"net/http"	
	"github.com/paulsumit5555/gowebdevlopment/mvc/controller"
	"html/template"
	"fmt"		
)

var tpl *template.Template

func init(){
	fmt.Println("inside init")
tpl = template.Must(template.ParseGlob("view/*"))
}

func main(){
	router := httprouter.New()
	uc := controller.GetUserController(tpl)
	fmt.Println(uc)
	router.GET("/home",uc.Home)
	http.ListenAndServe("localhost:8080", router)

}

/*func Home(w http.ResponseWriter, _ *http.Request, _ httprouter.Params){
	fmt.Println(tpl)
	tpl.ExecuteTemplate(w,"index.gohtml",nil)
}*/