package models

import (
	"crypto/md5"
	"encoding/hex"
	"hash/adler32"
	"log"
	"strings"
	"time"
)

// Account is a representation of a person who own a wallet and it's personal data
type Account struct {
	Name        string `bson:"name" json:"name"`
	Surname     string `bson:"surname" json:"surname"`
	Mail        string `bson:"mail" json:"mail"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`

	LoginID           uint32 `bson:"loginid" json:"loginid"`
	SocialInsuranceID string `bson:"socialinsuranceid" json:"socialinsuranceid"`
	PasswordHash      string `bson:"passwordhash" json:"passwordhash"`

	RegistrationDate time.Time `bson:"registrationdate" json:"registrationdate"`

	Wallets []Wallet
}

// NewAccount creates a new account assigned to a specified person
func NewAccount(name string, surname string, mail string, phonenumber string, socialinsuranceid string, passwordhash string) *Account {
	a := new(Account)
	a.Name = name
	a.Surname = surname
	a.Mail = mail
	a.PhoneNumber = phonenumber
	a.SocialInsuranceID = socialinsuranceid
	a.PasswordHash = passwordhash
	a.RegistrationDate = time.Now()
	return a
}

// AssignNewWallet checks if given wallet can be appended to account's wallet and appends it.
func (account *Account) AssignNewWallet(w Wallet) {
	if account.Wallets == nil {
		account.Wallets = make([]Wallet, 0)
		account.Wallets = append(account.Wallets, w)
		return
	}

	for i := 0; i < len(account.Wallets); i++ {
		if account.Wallets[i].Currency == w.Currency {
			log.Fatal("account with the given currency already exists")
			return
		}
	}

	account.Wallets = append(account.Wallets, w)

}

// GenerateLoginID generates a LoginID for the given account
func (account *Account) GenerateLoginID() {
	if account.LoginID != 0 {
		log.Fatal("This account's LoginID is already generated")
		return
	}

	concatenatedString := stringConcatenation(
		account.Name,
		account.Surname,
		account.SocialInsuranceID,
		account.RegistrationDate.Format(("20060102150405")))

	md5Hash := md5.Sum([]byte(concatenatedString))
	md5HashToString := hex.EncodeToString(md5Hash[:])
	adler32Hash := adler32.Checksum([]byte(md5HashToString))

	account.LoginID = adler32Hash
}

func stringConcatenation(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
