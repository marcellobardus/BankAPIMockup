package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/config"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/forms"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/responses"

	"net/http"
	"time"
)

func userLogin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	// Get payload data

	var login forms.UserLoginForm

	if err := json.NewDecoder(req.Body).Decode(&login); err != nil {
		response := responses.NewLoginResponse(true, nil, "Invalid request payload")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	// Check if login and password are valid

	account, err := database.GetAccountByLoginID(login.LoginID)

	if err != nil {
		response := responses.NewLoginResponse(true, nil, "Wrong LoginID there is no account with the given one")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	if login.PasswordHash != account.PasswordHash {
		response := responses.NewLoginResponse(true, nil, "Wrong password")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	// 2FA

	authenticated, err := account.OTP.Authenticate(login.OTP)
	if err != nil && err.Error() != "invalid code" {
		response := responses.NewLoginResponse(true, nil, "A fatal error occured")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	if !authenticated {
		response := responses.NewLoginResponse(true, nil, "OTP is wrong")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	// Create new authorization

	authorization := models.NewAuthorization([]string{"/sendTransaction"}, time.Now().Add(config.UserSessionExpirationMinutes), account)

	err = database.InsertAuthorization(authorization)

	if err != nil {
		response := responses.NewLoginResponse(true, nil, "A fatal error occured")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	// Return authorization token

	response := responses.NewLoginResponse(false, &authorization.Token, "Success")

	responseJSON, _ := json.Marshal(response)

	w.Write(responseJSON)
	http.Error(w, "Success", 200)
	return
}
