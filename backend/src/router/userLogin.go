package router

import (
	"bytes"
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/config"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"log"
	"net/http"
	"time"
)

func userLogin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	// Set response buffer

	var buffer bytes.Buffer

	// Get payload data

	var login models.UserLoginForm

	if err := json.NewDecoder(req.Body).Decode(&login); err != nil {
		buffer.WriteString(`{err: true, message: "Invalid request payload"}`)
		json.NewEncoder(w).Encode(buffer.String())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if login and password are valid

	account, err := database.GetAccountByLoginID(login.LoginID)

	if err != nil {
		buffer.WriteString(`{err: true, message: "Wrong LoginID there is no account with the given one"}`)
		json.NewEncoder(w).Encode(buffer.String())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if login.PasswordHash != account.PasswordHash {
		buffer.WriteString(`{err: true, message: "Wrong password"}`)
		json.NewEncoder(w).Encode(buffer.String())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2FA

	authenticated, err := account.OTP.Authenticate(login.OTP)
	if err != nil && err.Error() != "invalid code" {
		buffer.WriteString(`{err: true, message: "A fatal error occured"}`)
		json.NewEncoder(w).Encode(buffer.String())
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	if !authenticated {
		buffer.WriteString(`{err: true, message: "OTP is wrong"}`)
		json.NewEncoder(w).Encode(buffer.String())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create new authorization

	authorization := models.NewAuthorization([]string{"/sendTransaction"}, time.Now().Add(config.UserSessionExpirationMinutes), account)

	err = database.InsertAuthorization(authorization)

	if err != nil {
		buffer.WriteString(`{err: true, message: "A fatal error occured"}`)
		json.NewEncoder(w).Encode(buffer.String())
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	// Return authorization token

	response := models.NewLoginResponse(false, authorization.Token, "Success")

	responseJSON, err := json.Marshal(response)

	w.Write(responseJSON)
	http.Error(w, "Success", 200)
	return
}
