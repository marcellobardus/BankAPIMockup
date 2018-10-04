package router

import (
	"github.com/gorilla/mux"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/dao"
)

var database dao.BankMockupDAO

// GetRouter returns a router which just needs to be run
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	database.Database = "bankmockupdb"
	database.Server = "localhost"

	router.HandleFunc("/assignNewWalletToUser", assingNewWalletToUser).Methods("POST")
	router.HandleFunc("/createNewUser", registerUser).Methods("POST")
	router.HandleFunc("/sendTransaction", sendTransaction).Methods("POST")
	router.HandleFunc("/deposit", deposit).Methods("POST")
	router.HandleFunc("/revertTransaction", reverseTransaction).Methods("POST")
	router.HandleFunc("/walletHistory", getWalletHistory).Methods("GET")

	return router
}
