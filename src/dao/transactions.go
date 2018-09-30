package dao

import (
	"errors"
	"github.com/spaghettiCoderIT/BankAPIMockup/src/models"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// InsertTransaction inserts a new struct Transaction {...} into mongoDB collection: transactions
func (dao *BankMockupDAO) InsertTransaction(transaction *models.Transaction) error {
	if transaction.TransactionHash == "" {
		return errors.New("Could not insert the transaction because it's hash is not set")
	}

	existingTransaction, selectionError := dao.GetTransactionByHash(transaction.TransactionHash)
	if selectionError != nil && selectionError.Error() != "not found" {
		log.Println(selectionError.Error())
		panic("An error occured while inserting into db: ")
	}

	if existingTransaction != nil {
		return errors.New("Could not insert the new transaction, because one with the same hash already exists. Try to execute it later")
	}

	err := db.C(TransactionsCollection).Insert(transaction)
	return err
}

// GetTransactionByHash selects a specified transaction from collection: transactions
func (dao *BankMockupDAO) GetTransactionByHash(hash string) (*models.Transaction, error) {
	var transaction *models.Transaction
	err := db.C(TransactionsCollection).Find(bson.M{"transactionhash": hash}).One(&transaction)
	return transaction, err
}

func (dao *BankMockupDAO) GetTransactionsByRecipientIBAN(iban string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := db.C(TransactionsCollection).Find(bson.M{}).All(&transactions)
	return transactions, err
}
