package main

import(
	"net/http"
	"io"
	"fmt"
)





func main(){

	http.HandleFunc("/river",river)
	//http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("./flower"))))
	http.Handle("/",http.FileServer(http.Dir("./river")))
	http.ListenAndServe(":8080",nil)
}

func river(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL);
	w.Header().Set("Content-Type","Text/html;charset=utf-8")
	io.WriteString(w,`<img src="river.png">`)
	//io.WriteString(w,`<img src="/resources/river.png">`)
}

