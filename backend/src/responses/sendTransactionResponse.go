package responses

type SendTransactionResponse struct {
	Err             bool    `json:"err"`
	TransactionHash *string `json:"transactionhash"`
	Message         string  `json:"message"`
	ErrorCode       uint8   `json:"errorcode"`
}

func NewSendTransactionResponse(err bool, transactionHash *string, message string) *SendTransactionResponse {
	res := new(SendTransactionResponse)
	res.Err = err
	res.TransactionHash = transactionHash
	res.Message = message
	return res
}
