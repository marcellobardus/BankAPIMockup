package dao

import (
	"log"

	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BankMockupDAO struct defines mongoDB storage location
type BankMockupDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// AccountsCollection cointains the name of accounts collection name in mongoDB
	AccountsCollection = "accounts"
)

// ConnectToDatabase function establishes a connection to MongoDB databse
func (dao *BankMockupDAO) ConnectToDatabase() {
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(dao.Database)
}

// InsertAccount inserts a new struct Account {...} into mongoDB
func (dao *BankMockupDAO) InsertAccount(account *models.Account) error {
	err := db.C(AccountsCollection).Insert(account)
	return err
}

// Delete deletes an account from the database
func (dao *BankMockupDAO) Delete(account *models.Account) error {
	err := db.C(AccountsCollection).Remove(&account)
	return err
}

// UpdateByInsuranceID updates an index by it's social security number or PESEL etc...
func (dao *BankMockupDAO) UpdateByInsuranceID(id string, account *models.Account) error {
	err := db.C(AccountsCollection).Update(bson.M{"socialinsuranceid": id}, &account)
	return err
}

// GetAllAccounts return a list of all existing accounts in the database
func (dao *BankMockupDAO) GetAllAccounts() (*[]models.Account, error) {
	var accounts []models.Account
	err := db.C(AccountsCollection).Find(bson.M{}).All(&accounts)
	return &accounts, err
}

// GetAccountBySocialInsuranceID selects a customer by it's social security number or PESEL etc...
func (dao *BankMockupDAO) GetAccountBySocialInsuranceID(id string) (*models.Account, error) {
	var account models.Account
	err := db.C(AccountsCollection).Find(bson.M{"socialinsuranceid": id}).One(&account)
	return &account, err
}
