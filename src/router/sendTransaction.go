package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"net/http"
)

func sendTransaction(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	var transactionForm models.TransactionSendForm

	if err := json.NewDecoder(req.Body).Decode(&transactionForm); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if transactionForm.Amount <= 0 {
		http.Error(w, "Transaction amount must be greater than 0", http.StatusBadRequest)
		return
	}

	if transactionForm.RecipientIBAN == transactionForm.SenderIBAN {
		http.Error(w, "sender and recipient IBANs cannot be equal", http.StatusBadRequest)
		return
	}

	senderAccount, recErr := database.GetAccountBySocialInsuranceID(transactionForm.SenderSocialInsuranceID)
	recipientAccount, sendErr := database.GetAccountByWalletIBAN(transactionForm.Currency, transactionForm.RecipientIBAN)

	if recErr != nil && recErr.Error() == "not found" {
		http.Error(w, "the given sender IBAN is uncorrect", http.StatusBadRequest)
	}

	if sendErr != nil && sendErr.Error() == "not found" {
		http.Error(w, "the given recipient IBAN is uncorrect", http.StatusBadRequest)
	}

	if len(senderAccount.Wallets) == 0 || len(recipientAccount.Wallets) == 0 {
		http.Error(w, "One of the given accounts or both don't have any wallet assigned", http.StatusInternalServerError)
		return
	}

	transaction := models.NewTransaction(
		senderAccount,
		recipientAccount,
		transactionForm.Amount,
		transactionForm.Currency)

	transaction.SetFee()
	transaction.Realise()

	if err := database.InsertTransaction(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := database.UpdateAccountByInsuranceID(senderAccount.SocialInsuranceID, senderAccount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := database.UpdateAccountByInsuranceID(recipientAccount.SocialInsuranceID, recipientAccount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "Success", 200)
	return
}
