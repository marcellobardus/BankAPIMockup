package responses

type LoginResponse struct {
	err       bool
	token     *string
	message   string
	errorCode uint8
}

func NewLoginResponse(err bool, token *string, message string) *LoginResponse {
	res := new(LoginResponse)
	res.err = err
	res.token = token
	res.message = message
	return res
}
