package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/utils"
	"log"
	"net/http"
)

func sendTransaction(w http.ResponseWriter, req *http.Request) {
	const thisEndpoint = "/sendTransaction"

	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	// Get request data

	var transactionForm models.TransactionSendForm

	if err := json.NewDecoder(req.Body).Decode(&transactionForm); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if transactionForm.Amount <= 0 {
		http.Error(w, "Transaction amount must be greater than 0", http.StatusBadRequest)
		return
	}

	// Authorize token

	senderAuthorization, authErr := database.GetAuthorizationByToken(transactionForm.AuthorizationToken)

	if authErr != nil {
		http.Error(w, "Your token is obsolete you need to relogin", http.StatusForbidden)
		return
	}

	if !utils.IsInArray(thisEndpoint, senderAuthorization.Endpoints) {
		http.Error(w, "Your token is not authorized to call this endpoint", http.StatusForbidden)
		return
	}

	// Identify sender account by its token

	senderAccount := senderAuthorization.AuthorizedAccount

	// Check transfer conditions

	if _, exists := senderAccount.Wallets[transactionForm.Currency]; !exists {
		http.Error(w, "You're trying to withdraw a currency which you haven't a wallet assigned", http.StatusBadRequest)
		return
	}

	// Identify recipient account by its IBAN

	recipientAccount, recErr := database.GetAccountByWalletIBAN(transactionForm.Currency, transactionForm.RecipientIBAN)

	if recErr != nil {
		http.Error(w, "the given sender IBAN is uncorrect: "+recErr.Error(), http.StatusBadRequest)
		return
	}

	if len(senderAccount.Wallets) == 0 || len(recipientAccount.Wallets) == 0 {
		http.Error(w, "One of the given accounts or both don't have any wallet assigned", http.StatusInternalServerError)
		return
	}

	// Create a transaction

	transaction := models.NewTransaction(
		senderAccount,
		recipientAccount,
		transactionForm.Amount,
		transactionForm.Currency)

	transaction.SetFee()

	// Check if OTP is valid

	authenticated, err := senderAccount.OTP.Authenticate(transactionForm.OTP)
	if err != nil && err.Error() != "invalid code" {
		http.Error(w, "A fatal error occured", http.StatusInternalServerError)
		log.Panic(err.Error())
	}

	if !authenticated {
		http.Error(w, "OTP is wrong", http.StatusBadRequest)
		return
	}

	// Realise transaction and update accounts in db

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
