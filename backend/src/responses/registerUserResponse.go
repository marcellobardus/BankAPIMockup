package responses

type RegisterUserResponse struct {
	err                 bool
	authenticatorSecret *string
	loginID             *uint32
	message             string
	errorCode           uint16
}

func NewRegisterUserResponse(err bool, authenticatorSecret *string, loginID *uint32, message string) *RegisterUserResponse {
	res := new(RegisterUserResponse)
	res.err = err
	res.authenticatorSecret = authenticatorSecret
	res.loginID = loginID
	res.message = message
	return res
}
