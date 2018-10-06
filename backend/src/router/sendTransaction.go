package router

import (
	"encoding/json"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/forms"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/responses"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/utils"
	"log"
	"net/http"
)

func sendTransaction(w http.ResponseWriter, req *http.Request) {
	const thisEndpoint = "/sendTransaction"

	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	// Get request data

	var transactionForm forms.TransactionSendForm

	if err := json.NewDecoder(req.Body).Decode(&transactionForm); err != nil {
		response := responses.NewSendTransactionResponse(true, nil, "Invalid request payload")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	if transactionForm.Amount <= 0 {
		response := responses.NewSendTransactionResponse(true, nil, "Transaction amount must be greater than 0")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	// Authorize token

	senderAuthorization, authErr := database.GetAuthorizationByToken(transactionForm.AuthorizationToken)

	if authErr != nil {
		response := responses.NewSendTransactionResponse(true, nil, "Your token is obsolete you need to relogin")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	if !utils.IsInArray(thisEndpoint, senderAuthorization.Endpoints) {
		response := responses.NewSendTransactionResponse(true, nil, "Your token is not authorized to call this endpoint")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	// Identify sender account by its token

	senderAccount := senderAuthorization.AuthorizedAccount

	// Check transfer conditions

	if _, exists := senderAccount.Wallets[transactionForm.Currency]; !exists {
		response := responses.NewSendTransactionResponse(true, nil, "You're trying to withdraw a currency which you haven't a wallet assigned")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	// Identify recipient account by its IBAN

	recipientAccount, recErr := database.GetAccountByWalletIBAN(transactionForm.Currency, transactionForm.RecipientIBAN)

	if recErr != nil {
		response := responses.NewSendTransactionResponse(true, nil, "the given sender IBAN is uncorrect")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	if len(senderAccount.Wallets) == 0 || len(recipientAccount.Wallets) == 0 {
		response := responses.NewSendTransactionResponse(true, nil, "One of the given accounts or both don't have any wallet assigned")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
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
		response := responses.NewSendTransactionResponse(true, nil, "A fatal error occured")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		log.Panic(err.Error())
	}

	if !authenticated {
		response := responses.NewSendTransactionResponse(true, nil, "OTP is wrong")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	// Realise transaction and update accounts in db

	transaction.Realise()

	if err := database.InsertTransaction(transaction); err != nil {
		response := responses.NewSendTransactionResponse(true, nil, "A fatal error occured")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	if err := database.UpdateAccountByInsuranceID(senderAccount.SocialInsuranceID, senderAccount); err != nil {
		response := responses.NewSendTransactionResponse(true, nil, "A fatal error occured")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	if err := database.UpdateAccountByInsuranceID(recipientAccount.SocialInsuranceID, recipientAccount); err != nil {
		response := responses.NewSendTransactionResponse(true, nil, "A fatal error occured")
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
		return
	}

	response := responses.NewSendTransactionResponse(false, &transaction.TransactionHash, "Success")
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
	return
}
