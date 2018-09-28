package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"

	// "math/big"      TODO change balances type from float64 to BigFloat
	"time"
)

// Wallet represent a bank account not assigned to anyone, it's created to let move wallets between customers
type Wallet struct {
	currency    string
	bankName    string
	bankCountry string
	balance     float64
	IBAN        uint32
}

// Account is a representation of a person who own a wallet and it's personal data
type account struct {
	name    string
	surname string
	mail    string

	loginID      uint32
	passwordHash string

	registrationTime time.Time

	wallets []Wallet
}

// RegistrationData this data is passed in the POST requets
type registrationData struct {
	name         string `bson:"name" json:"name"`
	surname      string `bson:"surname" json:"surname"`
	mail         string `bson:"mail" json:"mail"`
	passwordHash string `bson:"passwordHash" json:"passwordHash"`

	// First bank Account Data

	bankName    string `bson:"bankName" json:"bankName"`
	bankCountry string `bson:"bankCountry" json:"bankCountry"`
	currency    string `bson:"currency" json:"currency"`
}

func createNewUser(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var registerBuffer registrationData
	if err := json.NewDecoder(req.Body).Decode(&registerBuffer); err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	params := mux.Vars(req)
	_ = json.NewDecoder(req.Body).Decode(&registerBuffer)

	if len(params) == 0 {
		http.Error(w, "params are empty", 400)
		return
	}

	s := session.Copy()
	defer s.Close()
	c := s.DB("bank_mockup").C("accounts")

	var newUser account

	newUser.name = params["name"]
	newUser.surname = params["surname"]
	newUser.mail = params["mail"]
	newUser.loginID = generateLoginID(params["name"], params["surname"], params["mail"], time.Now())
	newUser.passwordHash = params["passwordHash"]
	newUser.registrationTime = time.Now()
	newUser.wallets = make([]Wallet, 0)
	newUser.wallets = append(newUser.wallets, generateWallet(params["currency"], params["bankName"], params["bankCountry"]))

	dbErr := c.Insert(&newUser)
	if dbErr != nil {
		if mgo.IsDup(dbErr) {
			log.Println("User already exists")
			return
		}
	}

}

func insertIntoDatabase() {

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
