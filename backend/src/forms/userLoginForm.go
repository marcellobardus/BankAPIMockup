package forms

type UserLoginForm struct {
	LoginID      uint32 `bson:"loginid" json:"loginid"`
	PasswordHash string `bson:"passwordhash" json:"passwordhash"`
	OTP          string `bson:"otp" json:"otp"`
}
