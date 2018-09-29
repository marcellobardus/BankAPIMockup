package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"

	"net/http"
)

func assingNewWalletToUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	var wallet *models.Wallet

	if err := json.NewDecoder(req.Body).Decode(&wallet); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	wallet.SetIBAN()
	wallet.ResetBalance()

	if err := database.InsertWallet(wallet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
