package responses

type AssignNewWalletToUSerResponse struct {
	Err        bool    `json:"err"`
	WalletHash *string `json:"wallethash"`
	Message    string  `json:"message"`
	ErrorCode  uint8   `json:"errorcode"`
}

func NewAssignNewWalletToUSerResponse(err bool, walletHash *string, message string) *AssignNewWalletToUSerResponse {
	res := new(AssignNewWalletToUSerResponse)
	res.Err = err
	res.WalletHash = walletHash
	res.Message = message
	return res
}
