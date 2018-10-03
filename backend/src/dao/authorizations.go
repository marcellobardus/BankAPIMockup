package dao

import (
	"errors"
	"github.com/spaghettiCoderIT/BankAPIMockup/backend/src/models"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

func (dao *BankMockupDAO) InsertAuthorization(authorization *models.Authorization) error {
	existingAuthorization, selectionErr := dao.GetAuthorizationByToken(authorization.Token)

	if selectionErr != nil && selectionErr.Error() != "not found" {
		log.Println(selectionErr.Error())
		panic("An error occured while inserting into db: ")
	}

	if existingAuthorization != nil {
		return errors.New("Could not insert the new authorization, because one with the same token already exists")
	}

	if authorization.Expiration.Unix() <= time.Now().Unix() {
		return errors.New("Could not insert a new authorization into db because it's already obsolete")
	}

	err := db.C(AccountsCollection).Insert(authorization)
	return err
}

func (dao *BankMockupDAO) GetAuthorizationByToken(token string) (*models.Authorization, error) {
	var authorization *models.Authorization
	err := db.C(AuthorizationsCollection).Find(bson.M{"token": token}).One(&authorization)

	if authorization.Expiration.Unix() <= time.Now().Unix() {
		dao.DeleteAuthorization(authorization)
		return nil, errors.New("Authorization with the given token has been automaticly deleted because it's token is already obsolete")
	}

	return authorization, err
}

func (dao *BankMockupDAO) DeleteAuthorizationByToken(token string) error {
	authorization, err := dao.GetAuthorizationByToken(token)
	if err != nil {
		return err
	}

	err = db.C(AuthorizationsCollection).Remove(authorization)
	return err
}

func (dao *BankMockupDAO) DeleteAuthorization(authorization *models.Authorization) error {
	err := db.C(AuthorizationsCollection).Remove(authorization)
	return err
}
