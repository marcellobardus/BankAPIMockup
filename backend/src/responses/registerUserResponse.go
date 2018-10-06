package responses

type RegisterUserResponse struct {
	Err                 bool    `json:"err"`
	AuthenticatorSecret *string `json:"authenticatorsecret"`
	LoginID             *uint32 `json:"loginid"`
	Message             string  `json:"message"`
	ErrorCode           uint16  `json:"errorcode"`
}

func NewRegisterUserResponse(err bool, authenticatorSecret *string, loginID *uint32, message string) *RegisterUserResponse {
	res := new(RegisterUserResponse)
	res.Err = err
	res.AuthenticatorSecret = authenticatorSecret
	res.LoginID = loginID
	res.Message = message
	return res
}
