package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/forms"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/responses"
	"net/http"
)

func registerUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	var registration forms.AccountCreationForm

	if err := json.NewDecoder(req.Body).Decode(&registration); err != nil {
		response := responses.NewRegisterUserResponse(true, nil, nil, "Invalid request payload")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
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
		response := responses.NewRegisterUserResponse(true, nil, nil, "A fatal error occured")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	response := responses.NewRegisterUserResponse(false, &account.OTP.Secret, &account.LoginID, "A fatal error occured")
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
	return
}
