package router

import (
	"encoding/json"
	"fmt"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func createNewUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	var registration models.RegisterForm

	if err := json.NewDecoder(req.Body).Decode(&registration); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	account := models.NewAccount(
		registration.Name,
		registration.Surname,
		registration.Mail,
		registration.PhoneNumber,
		registration.SocialInsuranceID,
		registration.PasswordHash)

	if err := database.InsertAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Success")
}
