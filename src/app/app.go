package app

import (
	"fmt"
	"net/http"

	"github.com/spaghettiCoderIT/BankAPIMockup/src/dao"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/router"
)

var Database dao.BankMockupDAO

// Run start the entire app
func Run() {
	Database.Database = "bankmockupdb"
	Database.Server = "0.0.0.0"
	Database.ConnectToDatabase()

	router := router.GetRouter()
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", router)
}
