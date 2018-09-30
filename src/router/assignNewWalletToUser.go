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

	if err := wallet.SetIBAN(); err != nil {
		message := "Error while generating IBAN: " + err.Error()
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	wallet.SetHash()

	ownerAccount, err := database.GetAccountBySocialInsuranceID(wallet.OwnerSocialInsuranceID)

	if err != nil && err.Error() != "not found" {
		http.Error(w, "Account with wallets owner social insurance id does not exists", http.StatusInternalServerError)
		return
	}

	ownerAccount.Wallets[wallet.Currency] = wallet

	if err := database.UpdateAccountByInsuranceID(wallet.OwnerSocialInsuranceID, ownerAccount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, "Success", 200)
	return
}
