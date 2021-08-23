package schema

type trans struct {
	Deposit_m method `json:"deposit_m"`
	Remit_m   method `json:"deposit_m"`
}

type method struct {
	Description string   `json:"description"`
	Choices     []string `json:"choices"`
}
type infoasset struct {
	Enabled             bool                   `json:"enabled"`
	Max_amount          uint64                 `json:"max_amount"`
	Sender_sep12_type   string                 `json:"sender_sep12_type"`
	Receiver_sep12_type string                 `json:"receiver_sep12_type"`
	Fields              map[string]interface{} `json:"fields,omitempty"`
}

type info struct {
	Recieve map[string]infoasset `json:"recieve,omitempty"`
}

type transreq struct {
	Amount      uint64                 `json:"amount"`
	Asset_code  string                 `json:"asset_code"`
	Sender_id   string                 `json:"sender_id"`
	Receiver_id string                 `json:"receiver_id"`
	Fields      map[string]interface{} `json:"fields,omitempty"`
}

type transaction struct {
	Id                 string                 `json:"id"`
	Stellar_account_id string                 `json:"stellar_account_id"`
	Stellar_memo_type  string                 `json:"stellar_memo_type"`
	Stellar_memo       string                 `json:"stellar_memo"`
	Extra_info         map[string]interface{} `json:"extra_info,omitempty"`
	How                string                 `json:"how"`
}

type transinfo struct {
	Id                      string `json:"id"`
	Status                  string `json:"status"`
	Amount_in               string `json:"amount_in"`
	Amount_out              string `json:"amount_out"`
	Amount_fee              string `json:"amount_fee"`
	Stellar_account_id      string `json:"stellar_account_id"`
	Stellar_memo_type       string `json:"stellar_memo_type"`
	Stellar_memo            string `json:"stellar_memo"`
	Started_at              int64  `json:"timestamp"`
	Stellar_transaction_id  string `json:"stellar_transaction_id"`
	External_transaction_id string `json:"external_transaction_id"`
}
