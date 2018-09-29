package router

import (
	"github.com/gorilla/mux"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/dao"
)

var database dao.BankMockupDAO

// GetRouter returns a router which just needs to be run
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	database.Database = "bankmockupdb"
	database.Server = "localhost"
	router.HandleFunc("/assignNewWalletToUser", assingNewWalletToUser).Methods("POST")
	router.HandleFunc("/createNewUser", createNewUser).Methods("POST")
	return router
}
