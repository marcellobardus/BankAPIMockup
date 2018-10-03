package models

// TransactionSendForm
type TransactionSendForm struct {
	Amount                  int64  `bson:"amount" json:"amount"`
	Currency                string `bson:"currency" json:"currency"`
	SenderIBAN              string `bson:"senderiban" json:"senderiban"`
	RecipientIBAN           string `bson:"recipientiban" json:"recipientiban"`
	SenderSocialInsuranceID string `bson:"sendersocialinsuranceid" json:"sendersocialinsuranceid"`

	AuthorizationToken string   `bson:"token" json:"token"`
	Google2FACode      [6]uint8 `bson:"2facode" json:"2facode"`
}
