package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/forms"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/responses"

	"net/http"
	"strconv"
)

func assingNewWalletToUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	var data *forms.WalletCreationForm

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		response := responses.NewLoginResponse(true, nil, "Invalid request payload")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	wallet := models.NewWallet(
		data.Currency,
		data.BankName,
		data.BankCountry,
		data.OwnerSocialInsuranceID)

	if err := wallet.SetIBAN(); err != nil {
		response := responses.NewLoginResponse(true, nil, "Error while generating IBAN")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}
	wallet.SetHash()

	ownerAccount, err := database.GetAccountBySocialInsuranceID(wallet.OwnerSocialInsuranceID)

	if err != nil && err.Error() != "not found" {
		response := responses.NewLoginResponse(true, nil, "Account with wallets owner social insurance id does not exists")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	ownerAccount.Wallets[wallet.Currency] = wallet

	if err := database.UpdateAccountByInsuranceID(wallet.OwnerSocialInsuranceID, ownerAccount); err != nil {
		response := responses.NewLoginResponse(true, nil, "A fatal error occured")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	walletHash := strconv.Itoa(int(wallet.DataHash))

	response := responses.NewLoginResponse(false, &walletHash, "Success")
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
	return
}
