package dao

import (
	"errors"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// InsertWallet inserts a new struct Account {...} into mongoDB collection: wallets
func (dao *BankMockupDAO) InsertWallet(wallet *models.Wallet) error {
	existingWallet, selectionError := dao.GetWalletByDataHash(wallet.DataHash)
	if selectionError != nil && selectionError.Error() != "not found" {
		log.Println(selectionError.Error())
		panic("An error occured while inserting into db: ")
	}

	if existingWallet != nil {
		walletHash := string(wallet.DataHash)
		message := "Could not create a new wallet because, wallet with this data set already exists, existing wallet hash: " + walletHash
		return errors.New(message)
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

// GetWalletByIBAN selects a specified wallet from collection: wallets
func (dao *BankMockupDAO) GetWalletByIBAN(id string) (*models.Wallet, error) {
	var wallet *models.Wallet
	err := db.C(WalletsCollection).Find(bson.M{"iban": id}).One(&wallet)
	return wallet, err
}

// GetWalletByDataHash selects a specified wallet from collection: wallets
func (dao *BankMockupDAO) GetWalletByDataHash(hash int64) (*models.Wallet, error) {
	var wallet *models.Wallet
	err := db.C(WalletsCollection).Find(bson.M{"datahash": hash}).One(&wallet)
	return wallet, err
}

// GetAllWallets returns all wallets in collection wallets
func (dao *BankMockupDAO) GetAllWallets() ([]*models.Wallet, error) {
	var wallets []*models.Wallet
	err := db.C(WalletsCollection).Find(bson.M{}).All(&wallets)
	return wallets, err
}
