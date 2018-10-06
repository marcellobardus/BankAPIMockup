package responses

type AssignNewWalletToUSerResponse struct {
	err        bool
	walletHash *string
	message    string
	errorCode  uint8
}

func NewAssignNewWalletToUSerResponse(err bool, walletHash *string, message string) *AssignNewWalletToUSerResponse {
	res := new(AssignNewWalletToUSerResponse)
	res.err = err
	res.walletHash = walletHash
	res.message = message
	return res
}
