package forms

// TransactionSendForm
type TransactionSendForm struct {
	Amount        int64  `bson:"amount" json:"amount"`
	Currency      string `bson:"currency" json:"currency"`
	RecipientIBAN string `bson:"recipientiban" json:"recipientiban"`

	AuthorizationToken string `bson:"token" json:"token"`
	OTP                string `bson:"otp" json:"otp"`
}
