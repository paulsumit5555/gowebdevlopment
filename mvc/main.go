package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"	
	"github.com/paulsumit5555/gowebdevlopment/mvc/controller"
	"html/template"
	"fmt"		

)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("view/*"))
}

func main(){
	router := httprouter.New()
	uc := controller.GetUserController(tpl,getSession())
	fmt.Println(uc)
	router.GET("/home",uc.Home)
	router.GET("/add",uc.Add)
	router.POST("/submit",uc.Submit)
	http.ListenAndServe("localhost:8080", router)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		fmt.Println("error while creating connection ",err)
		panic(err)
	}
	return s
}