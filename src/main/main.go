package main

import (
	"fmt"
	"net/http"

	"github.com/spaghettiCoderIT/BankAPIMockup/src/router"
)

func main() {
	const a = 9
	router := router.GetRouter()
	fmt.Print("Listening on 3000")
	http.ListenAndServe(":3000", router)
}
