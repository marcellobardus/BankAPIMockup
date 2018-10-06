package responses

type SendTransactionResponse struct {
	err             bool
	transactionHash *string
	message         string
	errorCode       uint8
}

func NewSendTransactionResponse(err bool, transactionHash *string, message string) *SendTransactionResponse {
	res := new(SendTransactionResponse)
	res.err = err
	res.transactionHash = transactionHash
	res.message = message
	return res
}
