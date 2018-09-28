package router

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"

	// "math/big"      TODO change balances type from float64 to BigFloat
	"time"
)

// Wallet represent a bank account not assigned to anyone, it's created to let move wallets between people
type Wallet struct {
	currency    string
	bankName    string
	bankCountry string
	balance     float64
	IBAN        uint32
}

// Account is a representation of a person who own a wallet and it's personal data
type Account struct {
	name    string
	surname string
	mail    string

	loginID      uint32
	passwordHash string

	registrationTime time.Time

	wallets []Wallet
}

// RegistrationData this data is passed in the POST requets
type RegistrationData struct {
	name         string
	surname      string
	mail         string
	passwordHash string

	// First bank Account Data

	bankName    string
	bankCountry string
	currency    string
}

func createNewUser(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var registrationData RegistrationData
	decoderErr := decoder.Decode(&registrationData)
	if decoderErr != nil {
		panic(decoderErr)
	}

	name := registrationData.name
	surname := registrationData.surname
	mail := registrationData.mail
	passwordHash := registrationData.passwordHash

	// Registration account data

	bankName := registrationData.bankName
	bankCountry := registrationData.bankCountry
	currency := registrationData.currency

	s := session.Copy()
	defer s.Close()
	c := s.DB("bank_mockup").C("accounts")

	var newUser Account

	newUser.name = name
	newUser.surname = surname
	newUser.mail = mail
	newUser.loginID = generateLoginID(name, surname, mail, time.Now())
	newUser.passwordHash = passwordHash
	newUser.registrationTime = time.Now()
	newUser.wallets = make([]Wallet, 0)
	newUser.wallets = append(newUser.wallets, generateWallet(currency, bankName, bankCountry))

	dbErr := c.Insert(&newUser)
	if dbErr != nil {
		if mgo.IsDup(dbErr) {
			log.Println("User already exists")
			return
		}
	}
}

func InsertIntoDatabase() {

}

// TODO
func generateLoginID(name string, surname string, email string, birthdate time.Time) uint32 {
	return 0
}

// TODO
func generateWallet(currency string, bankName string, bankCountry string) Wallet {
	return Wallet{currency: currency, balance: 0, IBAN: generateIBAN(currency, bankName, bankCountry)}
}

// TODO
func generateIBAN(currency string, bankName string, bankCountry string) uint32 {
	return 1000000000
}

// TODO
func getIBANCountryCode(country string) string {
	return ""
}

// TODO
func getIBANByCurrency(currency string) string {
	return ""
}

// TODO
func getIBANByBankName(bankName string) string {
	return ""
}

func hashPassword(password string) string {
	passBytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
