package main

import (
	"fmt"
	"net/http"

	"github.com/spaghettiCoderIT/BankAPIMockup/src/router"
)

func main() {
	const a = 9
	router := router.GetRouter()
	http.Handle("/", router)
	fmt.Print(a)
}
