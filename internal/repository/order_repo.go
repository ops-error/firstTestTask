package repository

import (
	"context"
	"firstTestTask/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	dataBase *sqlx.DB
}

func NewOrderRepo(dataBase *sqlx.DB) *OrderRepo {
	return &OrderRepo{dataBase}
}

func (repo *OrderRepo) GetFullOrder(ctx context.Context, uid string) (*models.Order, error) {
	order := models.Order{}
	//var deliveryStruct models.Delivery
	//var paymentStruct models.Payment
	//var itemStruct models.Item

	queryOrder := `SELECT order_uid, track_number, entry, locale, internal_signature,
	customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
	FROM orders WHERE order_uid = $1`

	queryDelivery := `SELECT name, phone, zip, city, address, region, email
	FROM delivery WHERE order_uid = $1`

	queryPayment := `SELECT "transaction", request_id, currency, provider, amount, payment_dt,
	bank, delivery_cost, goods_total, custom_fee FROM payment WHERE "transaction" = $1`

	queryItems := `SELECT chrt_id, price, rid, name,
	sale, "size", total_price, nm_id, brand, status
	FROM items WHERE order_uid = $1`

	if err := repo.dataBase.GetContext(ctx, &order, queryOrder, uid); err != nil {
		return nil, err
	}

	if err := repo.dataBase.GetContext(ctx, &order.Delivery, queryDelivery, uid); err != nil {
		return nil, err
	}

	if err := repo.dataBase.GetContext(ctx, &order.Payment, queryPayment, uid); err != nil {
		return nil, err
	}

	var items []models.Item
	if err := repo.dataBase.SelectContext(ctx, &items, queryItems, uid); err != nil {
		return nil, err
	}
	order.Items = items
	return &order, nil
}
