package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"net/http"
)

func registerUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	var registration models.AccountCreationForm

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

	account.GenerateLoginID()
	account.SetOPT()

	if err := database.InsertAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Your 2FA secret key is: " + account.OTP.Secret)
	return
}
