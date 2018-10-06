package forms

// WalletCreationForm struct is required as interface when creating a new wallet via REST
type WalletCreationForm struct {
	Currency    string `bson:"currency" json:"currency"`
	BankName    string `bson:"bankname" json:"bankname"`
	BankCountry string `bson:"bankcountry" json:"bankcountry"`

	OwnerSocialInsuranceID string `bson:"ownersocialinsuranceid" json:"ownersocialinsuranceid"`

	AuthorizationToken string   `bson:"token" json:"token"`
	Google2FACode      [6]uint8 `bson:"2facode" json:"2facode"`
}
