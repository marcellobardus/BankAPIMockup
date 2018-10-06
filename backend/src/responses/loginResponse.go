package responses

type LoginResponse struct {
	Err       bool    `json:"err"`
	Token     *string `json:"token"`
	Message   string  `json:"message"`
	ErrorCode uint8   `json:"errorCode"`
}

func NewLoginResponse(err bool, token *string, message string) *LoginResponse {
	res := new(LoginResponse)
	res.Err = err
	res.Token = token
	res.Message = message
	return res
}
