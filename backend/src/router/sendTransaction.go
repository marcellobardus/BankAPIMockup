package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
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

	senderAccount, sendErr := database.GetAccountBySocialInsuranceID(transactionForm.SenderSocialInsuranceID)

	if _, exists := senderAccount.Wallets[transactionForm.Currency]; !exists {
		http.Error(w, "You're trying to withdraw a currency which you haven't a wallet assigned", http.StatusBadRequest)
		return
	}

	recipientAccount, recErr := database.GetAccountByWalletIBAN(transactionForm.Currency, transactionForm.RecipientIBAN)

	if recErr != nil {
		http.Error(w, "the given sender IBAN is uncorrect: "+recErr.Error(), http.StatusBadRequest)
		return
	}

	if sendErr != nil {
		http.Error(w, "the given recipient IBAN is uncorrect: "+sendErr.Error(), http.StatusBadRequest)
		return
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
