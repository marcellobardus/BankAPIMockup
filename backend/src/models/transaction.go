package models

import (
	"crypto/sha256"
	"errors"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/config"
	"time"
)

// Transaction defines transaction
type Transaction struct {
	TimeOfLeaving       time.Time `bson:"timeofleaving" json:"timeofleaving"`
	TimeOfComing        time.Time `bson:"timeofcoming" json:"timeofcoming"`
	Amount              int64     `bson:"amount" json:"amount"`
	TransactionCurrency string    `bson:"transactioncurrency" json:"transactioncurrency"`

	Sender    *Account `bson:"sender" json:"sender"`
	Recipient *Account `bson:"recipient" json:"recipient"`

	TransactionFee int64 `bson:"transactionfee" json:"transactionfee"`

	TransactionHash string `bson:"transactionhash" json:"transactionhash"`

	Status TransactionStatus `bson:"transactionstatus" json:"transactionstatus"`
}

// NewTransaction creates new transaction
func NewTransaction(sender *Account, recipient *Account, amount int64, currency string) *Transaction {
	t := new(Transaction)
	t.Amount = amount
	t.Sender = sender
	t.Recipient = recipient
	t.TransactionCurrency = currency
	t.Status = Unconfirmed
	return t
}

// SetFee calculates and sets the transaction fee
func (transaction *Transaction) SetFee() {
	transaction.TransactionFee = 1 // TODO
	transaction.Status = Pending
}

// Realise realises the transaction if fee is already set
func (transaction *Transaction) Realise() error {
	if transaction.TransactionFee == 0 {
		return errors.New("Transaction fee is not set yet")
	}

	transaction.Sender.Wallets[transaction.TransactionCurrency].DecreaseBalance(transaction.Amount)
	transaction.TimeOfLeaving = time.Now()
	transaction.Recipient.Wallets[transaction.TransactionCurrency].IncreaseBalance(transaction.Amount)
	transaction.TimeOfComing = transaction.TimeOfLeaving.Add(config.SessionTime)
	transaction.Status = Realised
	transaction.setTransactionHash()
	return nil
}

func (transaction *Transaction) setTransactionHash() {
	amount := string(transaction.Amount)
	sender := transaction.Sender.Wallets[transaction.TransactionCurrency].IBAN
	recipient := transaction.Recipient.Wallets[transaction.TransactionCurrency].IBAN
	currency := transaction.TransactionCurrency
	timeOfLeaving := transaction.TimeOfLeaving.String()
	timeOfComing := transaction.TimeOfComing.String()

	hashData := []byte(stringConcatenation(
		timeOfLeaving,
		timeOfComing,
		amount,
		sender,
		recipient,
		currency))

	hash := sha256.Sum256(hashData)

	transaction.TransactionHash = string(hash[:])
}

// TransactionStatus defines transaction status
type TransactionStatus int8

const (
	// Unconfirmed transaction cannot be send bcause its fee is undefined
	Unconfirmed TransactionStatus = 0
	// Pending => Transction is waiting for the nearest session
	Pending TransactionStatus = 1
	// Realised => Sender and Recipient wallets balances are updated
	Realised TransactionStatus = 2
	// Cancelled => Transaction is cancelled
	Cancelled TransactionStatus = 3
	// Frozen => Transaction is frozen
	Frozen TransactionStatus = 4
)
