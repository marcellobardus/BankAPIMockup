package router

import (
	"github.com/gorilla/mux"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/dao"
)

var database dao.BankMockupDAO

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	database.Database = "bankmockupdb"
	database.Server = "localhost"
	router.HandleFunc("/createNewUser", createNewUser).Methods("POST")
	return router
}
