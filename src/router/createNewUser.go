package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func createNewUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()
	var account models.Account
	if err := json.NewDecoder(req.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := database.InsertAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("succes"))
}

func hashPassword(password string) string {
	passBytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
