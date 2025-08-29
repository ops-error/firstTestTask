package models

type Payment struct {
	Transaction   string  `db:"transaction" json:"transaction"`
	Request_id    string  `db:"request_id" json:"request_id"`
	Currency      string  `db:"currency" json:"currency"`
	Provider      string  `db:"provider" json:"provider"`
	Amount        float64 `db:"amount" json:"amount"`
	Payment_dt    int64   `db:"payment_dt" json:"payment_dt"`
	Bank          string  `db:"bank" json:"bank"`
	Delivery_cost float64 `db:"delivery_cost" json:"delivery_cost"`
	Goods_total   float64 `db:"goods_total" json:"goods_total"`
	Custom_fee    float64 `db:"custom_fee" json:"custom_fee"`
}
