package controller

import (
	"fmt"
)

type UserController struct {
	 tpl *template.Template
}


func (uc UserController) home(){
	fmt.Println(tpl)
}