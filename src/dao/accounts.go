package dao

import (
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"gopkg.in/mgo.v2/bson"
)

// InsertAccount inserts a new struct Account {...} into mongoDB collection: accounts
func (dao *BankMockupDAO) InsertAccount(account *models.Account) error {
	err := db.C(AccountsCollection).Insert(account)
	return err
}

// DeleteAccount deletes an account from the database collection: accounts
func (dao *BankMockupDAO) DeleteAccount(account *models.Account) error {
	err := db.C(AccountsCollection).Remove(account)
	return err
}

// UpdateAccountByInsuranceID updates an index in mongoDB collection: accounts by it's social security number or PESEL etc...
func (dao *BankMockupDAO) UpdateAccountByInsuranceID(id string, account *models.Account) error {
	err := db.C(AccountsCollection).Update(bson.M{"socialinsuranceid": id}, account)
	return err
}

// GetAllAccounts return a list of all existing accounts in the database collection: accounts
func (dao *BankMockupDAO) GetAllAccounts() (*[]*models.Account, error) {
	var accounts []*models.Account
	err := db.C(AccountsCollection).Find(bson.M{}).All(accounts)
	return &accounts, err
}

// GetAccountBySocialInsuranceID selects a customer by it's social security number or PESEL etc...
func (dao *BankMockupDAO) GetAccountBySocialInsuranceID(id string) (*models.Account, error) {
	var account models.Account
	err := db.C(AccountsCollection).Find(bson.M{"socialinsuranceid": id}).One(account)
	return &account, err
}
