package main

import (
	"fmt"
	"net/http"

	"github.com/spaghettiCoderIT/BankAPIMockup/src/router"
)

func main() {
	router := router.GetRouter()
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", router)
}
