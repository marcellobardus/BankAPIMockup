package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"

	"net/http"
)

func assingNewWalletToUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	var data *models.WalletCreationForm

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	wallet := models.NewWallet(
		data.Currency,
		data.BankName,
		data.BankCountry,
		data.OwnerSocialInsuranceID)

	wallet.SetIBAN()
	wallet.ResetBalance()

	if err := database.InsertWallet(wallet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, "Success", 200)
	return
}
