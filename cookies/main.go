package main

import(
	"net/http"	
	"fmt"
	"log"
)



func main(){

	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.HandleFunc("/setmore",setmore)
	http.ListenAndServe(":8080",nil)
}

func set(w http.ResponseWriter, r *http.Request){

	http.SetCookie(w,&http.Cookie{
		Name : "name",
		Value: "sumit",
		Path:	"/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, r *http.Request){
	c1, err := r.Cookie("name")
	if err != nil{
		log.Println(err)
	}else{
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}

	c2, err := r.Cookie("SName")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", c2)
	}

	c3, err := r.Cookie("FName")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
	}

}

func setmore(w http.ResponseWriter, r *http.Request){

http.SetCookie(w, &http.Cookie{
	Name : "SName",
	Value : "Rahul",
})

http.SetCookie(w, &http.Cookie{
	Name : "FName",
	Value : "Rohan",
})
	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}