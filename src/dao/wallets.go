package dao

import (
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"gopkg.in/mgo.v2/bson"
)

// InsertWallet inserts a new struct Account {...} into mongoDB collection: wallets
func (dao *BankMockupDAO) InsertWallet(wallet *models.Wallet) error {
	err := db.C(WalletsCollection).Insert(wallet)
	return err
}

func (dao *BankMockupDAO) DeleteWallet(wallet *models.Wallet) error {
	err := db.C(WalletsCollection).Remove(wallet)
	return err
}

func (dao *BankMockupDAO) UpdateWalletByOwnerSocialInsuranceID(wallet *models.Wallet, account *models.Account) error {
	err := db.C(WalletsCollection).Update(bson.M{"ownersocialinsuranceid": account.SocialInsuranceID}, wallet)
	return err
}

func (dao *BankMockupDAO) GetWalletByOwnerSocialInsuranceID(id string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := db.C(WalletsCollection).Find(bson.M{"ownersocialinsuranceid": id}).One(wallet)
	return &wallet, err
}

func (dao *BankMockupDAO) GetAllWallets() (*[]*models.Wallet, error) {
	var wallets []*models.Wallet
	err := db.C(WalletsCollection).Find(bson.M{}).All(wallets)
	return &wallets, err
}
