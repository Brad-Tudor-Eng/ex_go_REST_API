package main

import (
	"io"
	"log"
	"net/http"

	"./controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	uc := controllers.NewUserController(getSession())
	r := httprouter.New()

	r.GET("/", index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.RemoveUser)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Welcome")
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://admin:admin123@ds157136.mlab.com:57136/go_lang_rest_api_example")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
