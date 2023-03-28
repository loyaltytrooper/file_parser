package models

import "time"

type Transaction struct {
	Date         time.Time `json:"txn_time"`
	TxnType      string    `json:"txn_type"`
	TxnId        string    `json:"txn_id"`
	TransferMode string    `json:"payment_mode"`
	Destination  string    `json:"destination"`
	Credit       float64   `json:"credit"`
	Debit        float64   `json:"debit"`
	FinalAmount  float64   `json:"final_amount"`
}

type Transactions struct {
	Txns []Transaction `json:"transactions"`
}

type FrontendStatements struct {
	UserStatements map[string][]Transaction `json:"user_statements"`
}
