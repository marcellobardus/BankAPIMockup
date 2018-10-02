package dao

import (
	"errors"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

// InsertAccount inserts a new struct Account {...} into mongoDB collection: accounts
func (dao *BankMockupDAO) InsertAccount(account *models.Account) error {
	existingAccount, selectionError := dao.GetAccountBySocialInsuranceID(account.SocialInsuranceID)
	if selectionError != nil && selectionError.Error() != "not found" {
		panic("An error occured while inserting into db: ")
	}

	if existingAccount != nil {
		return errors.New("Could not create a new account because, one with the given SocialInsuranceID already exists")
	}

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
func (dao *BankMockupDAO) GetAllAccounts() ([]*models.Account, error) {
	var accounts []*models.Account
	err := db.C(AccountsCollection).Find(bson.M{}).All(&accounts)
	return accounts, err
}

// GetAccountBySocialInsuranceID selects a customer by it's social security number or PESEL etc...
func (dao *BankMockupDAO) GetAccountBySocialInsuranceID(id string) (*models.Account, error) {
	var account *models.Account
	err := db.C(AccountsCollection).Find(bson.M{"socialinsuranceid": id}).One(&account)
	return account, err
}

// GetAccountByWalletIBAN returns an account which cointains a wallet with the given currency and IBAN
func (dao *BankMockupDAO) GetAccountByWalletIBAN(currency string, iban string) (*models.Account, error) {
	var account *models.Account
	field := "wallets." + strings.ToUpper(currency) + ".iban"
	err := db.C(AccountsCollection).Find(bson.M{field: iban}).One(&account)

	if account == nil {
		return account, errors.New("There is no account with the given IBAN, or the given currency is invalid for the given IBAN")
	}

	return account, err
}
