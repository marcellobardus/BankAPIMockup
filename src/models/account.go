package models

import "time"

// Account is a representation of a person who own a wallet and it's personal data
type Account struct {
	Name    string `bson:"name" json:"name"`
	Surname string `bson:"surname" json:"surname"`
	Mail    string `bson:"mail" json:"mail"`

	LoginID           uint32 `bson:"loginid" json:"loginid"`
	SocialInsuranceID uint32 `bson:"socialinsuranceid" json:"socialinsuranceid"`
	PasswordHash      string `bson:"passwordhash" json:"passwordhash"`

	RegistrationDate time.Time `bson:"registrationdate" json:"registrationdate"`

	Wallets []Wallet
}
