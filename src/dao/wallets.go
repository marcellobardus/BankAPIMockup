package dao

import (
	"errors"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"gopkg.in/mgo.v2/bson"
)

// InsertWallet inserts a new struct Account {...} into mongoDB collection: wallets
func (dao *BankMockupDAO) InsertWallet(wallet *models.Wallet) error {
	existingWallet, selectionError := dao.GetWalletByOwnerSocialInsuranceID(wallet.OwnerSocialInsuranceID)
	if selectionError != nil && selectionError.Error() != "not found" {
		panic("An error occured while inserting into db: ")
	}

	if existingWallet != nil {
		return errors.New("Could not create a new wallet because, one with the given OwnerSocialInsuranceID already exists")
	}

	err := db.C(WalletsCollection).Insert(wallet)
	return err
}

// DeleteWallet deletes a given wallet from collection: wallets
func (dao *BankMockupDAO) DeleteWallet(wallet *models.Wallet) error {
	err := db.C(WalletsCollection).Remove(wallet)
	return err
}

// UpdateWalletByOwnerSocialInsuranceID updates an index in mongoDB collection: accounts
func (dao *BankMockupDAO) UpdateWalletByOwnerSocialInsuranceID(wallet *models.Wallet, account *models.Account) error {
	err := db.C(WalletsCollection).Update(bson.M{"ownersocialinsuranceid": account.SocialInsuranceID}, wallet)
	return err
}

// GetWalletByOwnerSocialInsuranceID selects a specified wallet from collection: wallets
func (dao *BankMockupDAO) GetWalletByOwnerSocialInsuranceID(id string) (*models.Wallet, error) {
	var wallet *models.Wallet
	err := db.C(WalletsCollection).Find(bson.M{"ownersocialinsuranceid": id}).One(&wallet)
	return wallet, err
}

// GetAllWallets returns all wallets in collection wallets
func (dao *BankMockupDAO) GetAllWallets() ([]*models.Wallet, error) {
	var wallets []*models.Wallet
	err := db.C(WalletsCollection).Find(bson.M{}).All(&wallets)
	return wallets, err
}
