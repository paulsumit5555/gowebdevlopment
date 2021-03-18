package main

import(
	"fmt"
	"net/http"
)


type handler int


 func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request){

 	fmt.Fprintln(w, "this is my first Http Server example")
 }


func main(){
	var h handler
	http.ListenAndServe(":8080",h)

}