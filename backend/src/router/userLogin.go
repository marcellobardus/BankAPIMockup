package router

import (
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

	// Get payload data

	var login models.UserLoginForm

	if err := json.NewDecoder(req.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if login and password are valid

	account, err := database.GetAccountByLoginID(login.LoginID)

	if err != nil {
		http.Error(w, "Wrong LoginID there is no account with the given one", http.StatusBadRequest)
		return
	}

	if login.PasswordHash != account.PasswordHash {
		http.Error(w, "Wrong password", http.StatusBadRequest)
		return
	}

	// Create new authorization

	authorization := models.NewAuthorization([]string{"/sendTransaction"}, time.Now().Add(config.UserSessionExpirationMinutes), account)

	err = database.InsertAuthorization(authorization)

	if err != nil {
		http.Error(w, "A fatal error occured", http.StatusInternalServerError)
		log.Panic(err.Error())
	}

	// Return authorization token

	tokenJSON, err := json.Marshal(authorization.Token)

	if err != nil {
		http.Error(w, "A fatal error occured", http.StatusInternalServerError)
		log.Panic(err.Error())
	}

	w.Write(tokenJSON)
	http.Error(w, "Success", 200)
	return
}
