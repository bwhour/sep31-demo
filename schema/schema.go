package schema

type trans struct {
	Deposit_m Method `json:"deposit_m"`
	Remit_m   Method `json:"deposit_m"`
}

type Method struct {
	Description string   `json:"description"`
	Choices     []string `json:"choices"`
}
type Infoasset struct {
	Enabled             bool                   `json:"enabled"`
	Max_amount          uint64                 `json:"max_amount"`
	Sender_sep12_type   string                 `json:"sender_sep12_type"`
	Receiver_sep12_type string                 `json:"receiver_sep12_type"`
	Fields              map[string]interface{} `json:"fields,omitempty"`
}

type Info struct {
	Recieve map[string]Infoasset `json:"recieve,omitempty"`
}

type Transreq struct {
	Amount      uint64                 `json:"amount"`
	Asset_code  string                 `json:"asset_code"`
	Sender_id   string                 `json:"sender_id"`
	Receiver_id string                 `json:"receiver_id"`
	Fields      map[string]interface{} `json:"fields,omitempty"`
}

type Transaction struct {
	Id                 string                 `json:"id"`
	Stellar_account_id string                 `json:"stellar_account_id"`
	Stellar_memo_type  string                 `json:"stellar_memo_type"`
	Stellar_memo       string                 `json:"stellar_memo"`
	Extra_info         map[string]interface{} `json:"extra_info,omitempty"`
	How                string                 `json:"how,omitempty"`
}

type Transinfo struct {
	Id                    string `json:"id"`
	Status                string `json:"status"`
	AmountIn              string `json:"amount_in,omitempty"`
	AmountOut             string `json:"amount_out,omitempty"`
	AmountFee             string `json:"amount_fee,omitempty"`
	StellarAccountId      string `json:"stellar_account_id"`
	StellarMemoType       string `json:"stellar_memo_type"`
	StellarMemo           string `json:"stellar_memo"`
	StartedAt             string `json:"started_at,omitempty"`
	StellarTransactionId  string `json:"stellar_transaction_id,omitempty"`
	ExternalTransactionId string `json:"external_transaction_id,omitempty"`
}

type CustomerInfo struct {
	Id              string                 `json:"id"`
	Status          string                 `json:"status"`
	Message         string                 `json:"message,omitempty"`
	Fields          map[string]interface{} `json:"fields,omitempty"`
	Provided_fields map[string]interface{} `json:"provided_fields,omitempty"`
}

type CustmoerReq struct {
	Account  string `json:"id,omitempty"`
	Memo     string `json:"memo,omitempty"`
	MemoType string `json:"memo_type,omitempty"`
	Type     string `json:"type,omitempty"`
}