package forms

// AccountCreationForm struct is required as interface when creating a new user via REST
type AccountCreationForm struct {
	Name        string `bson:"name" json:"name"`
	Surname     string `bson:"surname" json:"surname"`
	Mail        string `bson:"mail" json:"mail"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`

	SocialInsuranceID string `bson:"socialinsuranceid" json:"socialinsuranceid"`
	PasswordHash      string `bson:"passwordhash" json:"passwordhash"`

	AuthorizationToken string `bson:"token" json:"token"`
	Google2FACode      string `bson:"2facode" json:"2facode"`
}
