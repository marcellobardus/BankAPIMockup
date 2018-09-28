package models

import (
	"log"
)

// Wallet represent a bank account not assigned to anyone, it's created to let move wallets between customers
type Wallet struct {
	Currency    string  `bson:"currency" json:"currency"`
	BankName    string  `bson:"bankname" json:"bankname"`
	BankCountry string  `bson:"bankcountry" json:"bankcountry"`
	IBAN        string  `bson:"iban" json:"iban"`
	balance     float64 `bson:"balance"`
}

// NewWallet creates a new Wallet args: currency ex. "USD", bankName ex. "AliorBank", bankCountry ex. "PL"
func NewWallet(currency string, bankName string, bankCountry string) *Wallet {
	w := new(Wallet)
	w.Currency = currency
	w.BankName = bankName
	w.BankCountry = bankCountry
	return w
}

// IncreaseBalance increases wallet balance by the given amount
func (wallet *Wallet) IncreaseBalance(amount float64) {
	if amount < 0 {
		log.Fatal("amount can not be less than 0")
		return
	}

	wallet.balance += amount
}

// DecreaseBalance decreases wallet balance by the given amount
func (wallet *Wallet) DecreaseBalance(amount float64) {
	if amount < 0 {
		log.Fatal("amount can not be less than 0")
		return
	}

	wallet.balance -= amount
}

// SetIBAN generates and sets a new IBAN for the given wallet, if IBAN is already set throws an error
func (wallet *Wallet) SetIBAN() {
	if wallet.IBAN != "" {
		log.Fatal("IBAN for this wallet is already set")
		return
	}

	wallet.IBAN = "NotARealIBAN" // TODO
}
