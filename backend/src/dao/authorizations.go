package dao

import (
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
)

func (dao *BankMockupDAO) InsertAuthorization(authorization *models.Authorization) error {
	err := db.C(AccountsCollection).Insert(authorization)
	return err
}
