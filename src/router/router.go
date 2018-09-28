package router

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetRouter() mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/createNewUser", createNewUser).Methods("POST")
	return *router
}
