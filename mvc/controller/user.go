package controller

import (
	"html/template"
	"github.com/julienschmidt/httprouter"
	"github.com/paulsumit5555/gowebdevlopment/mvc/model"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"fmt"
//	"io"
  //  "log"
    //"errors"
    //"strings"
)

type UserController struct {
	 tpl *template.Template
	 session *mgo.Session
}

func GetUserController(tpl *template.Template,session *mgo.Session) *UserController{
	return &UserController{tpl,session}
}


func (uc UserController) Home(w http.ResponseWriter, _ *http.Request, _ httprouter.Params){
	uc.tpl.ExecuteTemplate(w,"index.gohtml",nil)
}

func (uc UserController) Add(w http.ResponseWriter, _ *http.Request, _ httprouter.Params){
	uc.tpl.ExecuteTemplate(w,"addEmployee.gohtml",nil)
}

func (uc UserController) Submit(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	
	r.ParseForm()
	//json.NewDecoder(r.Body).Decode(&u)

/*	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
    if err != nil {
        var syntaxError *json.SyntaxError
        var unmarshalTypeError *json.UnmarshalTypeError

        switch {
        // Catch any syntax errors in the JSON and send an error message
        // which interpolates the location of the problem to make it
        // easier for the client to fix.
        case errors.As(err, &syntaxError):
            msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
            http.Error(w, msg, http.StatusBadRequest)

        // In some circumstances Decode() may also return an
        // io.ErrUnexpectedEOF error for syntax errors in the JSON. There
        // is an open issue regarding this at
        // https://github.com/golang/go/issues/25956.
        case errors.Is(err, io.ErrUnexpectedEOF):
            msg := fmt.Sprintf("Request body contains badly-formed JSON")
            http.Error(w, msg, http.StatusBadRequest)

        // Catch any type errors, like trying to assign a string in the
        // JSON request body to a int field in our Person struct. We can
        // interpolate the relevant field name and position into the error
        // message to make it easier for the client to fix.
        case errors.As(err, &unmarshalTypeError):
            msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
            http.Error(w, msg, http.StatusBadRequest)

        // Catch the error caused by extra unexpected fields in the request
        // body. We extract the field name from the error message and
        // interpolate it in our custom error message. There is an open
        // issue at https://github.com/golang/go/issues/29035 regarding
        // turning this into a sentinel error.
        case strings.HasPrefix(err.Error(), "json: unknown field "):
            fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
            msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
            http.Error(w, msg, http.StatusBadRequest)

        // An io.EOF error is returned by Decode() if the request body is
        // empty.
        case errors.Is(err, io.EOF):
            msg := "Request body must not be empty"
            http.Error(w, msg, http.StatusBadRequest)

        // Catch the error caused by the request body being too large. Again
        // there is an open issue regarding turning this into a sentinel
        // error at https://github.com/golang/go/issues/30715.
        case err.Error() == "http: request body too large":
            msg := "Request body must not be larger than 1MB"
            http.Error(w, msg, http.StatusRequestEntityTooLarge)

        // Otherwise default to logging the error and sending a 500 Internal
        // Server Error response.
        default:
            log.Println(err.Error())
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
	return
    }
*/

	//create bson id

u := models.User{r.Form["name"][0],r.Form["gender"][0],22,bson.NewObjectId()}
	fmt.Println("age ", r.Form["age"])
	fmt.Println("body ", u)


	dbs,dberr := uc.session.DatabaseNames()
	
		fmt.Println("dbs name", dbs)
	if dberr != nil{
		fmt.Println("dbnames error is", dberr)
	}

	// store the user in mongodb
	db  := uc.session.DB("usersdb")
fmt.Println("database ", db)
	c  := db.C("user")
	fmt.Println("collection ", c)
	err := c.Insert(u)
	fmt.Println("error ", err)
	
	if err!= nil{
		fmt.Println("error while inserting record ",err)
	}

	
		
		
	fmt.Println(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s\n", uj)
	//uc.tpl.ExecuteTemplate(w,"addEmployee.gohtml",nil)
}