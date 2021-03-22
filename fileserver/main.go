package main

import(
	"net/http"
	"io"
)





func main(){

	http.Handle("/",http.FileServer(http.Dir(".")))
	http.HandleFunc("/river",river)
	http.HandleFunc("/flower",flower)
	http.ListenAndServe(":8080",nil)
}

func river(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","Text/html;charset=utf-8")
	io.WriteString(w,`<img src="river.png">`)
}

func flower(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","Text/html;charset=utf-8")
	io.WriteString(w,`<img src="flower.png">`)
}