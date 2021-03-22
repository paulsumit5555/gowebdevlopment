package controller

import (
	"html/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

type UserController struct {
	 tpl *template.Template
}

func GetUserController(tpl *template.Template) *UserController{
	return &UserController{tpl}
}


func (uc UserController) Home(w http.ResponseWriter, _ *http.Request, _ httprouter.Params){
	fmt.Println(uc.tpl)
	uc.tpl.ExecuteTemplate(w,"index.gohtml",nil)
}