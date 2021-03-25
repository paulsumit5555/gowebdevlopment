package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/paulsumit5555/gowebdevlopment/mvc/controller"
	"gopkg.in/mgo.v2"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("view/*"))
}

func main() {
	router := httprouter.New()
	uc := controller.GetUserController(tpl, getSession())
	fmt.Println(uc)
	router.GET("/home", uc.Home)
	router.GET("/add", uc.Add)
	router.POST("/submit", uc.Submit)
	err := http.ListenAndServe("localhost:9090", router)
	if err != nil {
		fmt.Println(err)
	}

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		fmt.Println("error while creating connection ", err)
		panic(err)
	}
	return s
}
