package models

// TransactionSendForm
type TransactionSendForm struct {
	Amount                  int64  `bson:"amount" json:"amount"`
	Currency                string `bson:"currency" json:"currency"`
	SenderIBAN              string `bson:"senderiban" json:"senderiban"`
	RecipientIBAN           string `bson:"recipientiban" json:"recipientiban"`
	SenderSocialInsuranceID string `bson:"sendersocialinsuranceid" json:"sendersocialinsuranceid"`

	SenderPasswordHash string `bson:"senderpasswordhash" json:"senderpasswordhash"`
	SenderLoginID      uint32 `bson:"senderloginid" json:"senderloginid"`
	SenderPhoneNumber  string `bson:"senderphonenumber" json:"senderphonenumber"`
}
