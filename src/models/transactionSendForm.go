package models

type TransactionSendForm struct {
	Amount        int64  `bson:"amount" json:"amount"`
	SenderIBAN    string `bson:"senderiban" json:"senderiban"`
	RecipientIBAN string `bson:"recipientiban" json:"recipientiban"`

	SenderPasswordHash string `bson:"senderpasswordhash" json:"senderpasswordhash"`
	SenderLoginID      uint32 `bson:"senderloginid" json:"senderloginid"`
	SenderPhoneNumber  string `bson:"senderphonenumber" json:"senderphonenumber"`
}
