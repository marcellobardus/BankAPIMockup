package models

// Wallet represent a bank account not assigned to anyone, it's created to let move wallets between customers
type Wallet struct {
	Currency    string `bson:"currency" json:"currency"`
	BankName    string `bson:"bankname" json:"bankname"`
	BankCountry string `bson:"bankcountry" json:"bankcountry"`
	IBAN        uint64 `bson:"iban" json:"iban"`
	balance     float64
}
