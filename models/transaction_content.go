package models

import "time"

type Transaction struct {
	Date            time.Time `json:"txn_time"`
	Description     string    `json:"description"`
	ChequeReference string    `json:"cheque_reference"`
	Credit          float64   `json:"credit"`
	Debit           float64   `json:"debit"`
	FinalAmount     float64   `json:"final_amount"`
}

type Transactions struct {
	Txns []Transaction `json:"transactions"`
}
