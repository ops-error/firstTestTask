package models

import (
	"time"
)

type Order struct {
	Order_uid          int64     `db:"order_uid" json:"order_uid"`
	Track_number       string    `db:"order_uid" json:"track_number"`
	Entry              string    `db:"entry" json:"entry"`
	Delivery           Delivery  `db:"delivery" json:"delivery"`
	Payment            Payment   `db:"payment" json:"payment"`
	Items              []Item    `db:"items" json:"items"`
	Locale             string    `db:"locale" json:"locale"`
	Internal_signature string    `db:"internal_signature" json:"internal_signature"`
	Customer_id        string    `db:"customer_id" json:"customer_id"`
	Delivery_service   string    `db:"delivery_service" json:"delivery_service"`
	Shardkey           int64     `db:"shardkey" json:"shardkey"`
	Sm_id              int64     `db:"sm_id" json:"sm_id"`
	Date_created       time.Time `db:"date_created" json:"date_created"`
	Oof_shard          int64     `db:"oof_shard" json:"oof_shard"`
}
