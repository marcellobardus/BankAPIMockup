package models

import (
	"errors"
	"github.com/cespare/xxhash"
	"hash/adler32"
	"log"
	"strconv"
)

// Wallet represent a bank account not assigned to anyone, it's created to let move wallets between customers
type Wallet struct {
	Currency    string `bson:"currency" json:"currency"`
	BankName    string `bson:"bankname" json:"bankname"`
	BankCountry string `bson:"bankcountry" json:"bankcountry"`
	IBAN        string `bson:"iban" json:"iban"`
	Balance     int64  `bson:"balance" json:"balance"`

	OwnerSocialInsuranceID string `bson:"ownersocialinsuranceid" json:"ownersocialinsuranceid"`

	DataHash int64 `bson:"datahash" json:"datahash"`
}

// NewWallet creates a new Wallet args: currency ex. "USD", bankName ex. "AliorBank", bankCountry ex. "PL"
func NewWallet(currency string, bankName string, bankCountry string, ownerSocialInsuranceID string) *Wallet {
	w := new(Wallet)
	w.Currency = currency
	w.BankName = bankName
	w.BankCountry = bankCountry
	w.OwnerSocialInsuranceID = ownerSocialInsuranceID
	return w
}

// IncreaseBalance increases wallet balance by the given amount
func (wallet *Wallet) IncreaseBalance(amount int64) {
	if amount < 0 {
		log.Fatal("amount can not be less than 0")
		return
	}

	wallet.Balance += amount
}

// DecreaseBalance decreases wallet balance by the given amount
func (wallet *Wallet) DecreaseBalance(amount int64) {
	if amount < 0 {
		log.Fatal("amount can not be less than 0")
		return
	}

	wallet.Balance -= amount
}

// ResetBalance sets wallet balance to 0
func (wallet *Wallet) ResetBalance() {
	wallet.Balance = 0
}

// SetIBAN generates and sets a new IBAN for the given wallet, if IBAN is already set throws an error
func (wallet *Wallet) SetIBAN() error {
	if wallet.IBAN != "" {
		return errors.New("IBAN already set")
	}

	var codes map[string]string
	codes = make(map[string]string)
	codes["Poland"] = "PL"
	codes["UnitedStates"] = "US"
	codes["Germany"] = "DE"
	codes["GreatBritain"] = "GB"
	codes["Italy"] = "IT"

	if codes[wallet.BankCountry] == "" {
		return errors.New("The given country is not supported")
	}

	hash := adler32.Checksum([]byte(stringConcatenation(
		wallet.Currency,
		wallet.BankName,
		wallet.OwnerSocialInsuranceID)))

	hashString := strconv.FormatUint(uint64(hash)*100, 10)
	wallet.IBAN = codes[wallet.BankCountry] + string(hashString)
	return nil
}

// SetHash sets the wallet hash which proofs it's uniqueness
func (wallet *Wallet) SetHash() {
	data := stringConcatenation(
		wallet.OwnerSocialInsuranceID,
		wallet.IBAN,
		wallet.BankName,
		wallet.BankCountry,
		wallet.Currency)

	wallet.DataHash = int64(xxhash.Sum64String(data) / 2)
}
