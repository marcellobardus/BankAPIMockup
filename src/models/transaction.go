package models

import (
	"errors"
	"time"
)

// Transaction defines transaction
type Transaction struct {
	Time                time.Time `bson:"time" json:"time"`
	Amount              int64     `bson:"amount" json:"amount"`
	TransactionCurrency string    `bson:"transactioncurrency" json:"transactioncurrency"`

	Sender    *Account `bson:"sender" json:"sender"`
	Recipient *Account `bson:"recipient" json:"recipient"`

	TransactionFee int64 `bson:"transactionfee" json:"transactionfee"`

	TransactionHash string `bson:"transactionhash" json:"transactionhash"`
}

// NewTransaction creates new transaction
func (transaction *Transaction) NewTransaction(sender *Account, recipient *Account, amount int64) *Transaction {
	t := new(Transaction)
	t.Amount = amount
	t.Sender = sender
	t.Recipient = recipient
	return t
}

// SetFee calculates and sets the transaction fee
func (transaction *Transaction) SetFee() {
	transaction.TransactionFee = 1 // TODO
}

// Realise realises the transaction if fee is already set
func (transaction *Transaction) Realise() error {
	if transaction.TransactionFee == 0 {
		return errors.New("Transaction fee is not set yet")
	}

	transaction.Sender.Wallets[transaction.TransactionCurrency].DecreaseBalance(transaction.Amount)
	transaction.Recipient.Wallets[transaction.TransactionCurrency].IncreaseBalance(transaction.Amount)
	return nil
}
