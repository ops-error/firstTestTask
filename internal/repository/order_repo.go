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

const queryFullOrder = `
SELECT
	ord.order_uid, ord.track_number, ord.entry, ord.locale,
	ord.internal_signature, ord.customer_id, ord.delivery_service,
	ord.shardkey, ord.sm_id, ord.date_created, ord.oof_shard,
	del.name, del.phone, del.zip, del.city, del.address,
	del.region, del.email,
	pay.transaction, pay.request_id, pay.currency, pay.provider,
	pay.amount, pay.payment_dt, pay.bank,
	pay.delivery_cost, pay.goods_total, pay.custom_fee,
	itm.chrt_id, itm.price, itm.rid, itm.name, itm.sale,
	itm.size, itm.total_price, itm.nm_id, itm.brand, itm.status
FROM orders ord
LEFT JOIN delivery del ON del.order_uid = ord.order_uid
LEFT JOIN payment pay ON pay.transaction = ord.order_uid
LEFT JOIN items itm ON itm.order_uid = ord.order_uid
WHERE ord.order_uid = $1
ORDER BY itm.chrt_id
`

func (repo *OrderRepo) GetFullOrder(ctx context.Context, uid string) (*models.Order, error) {
	rows, err := repo.dataBase.QueryxContext(ctx, queryFullOrder, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		order models.Order
		items []models.Item
	)
	for rows.Next() {
		var (
			ord models.Order
			del models.Delivery
			pay models.Payment
			itm models.Item
		)

		if err := rows.Scan(
			&ord.Order_uid, &ord.Track_number, &ord.Entry, &ord.Locale,
			&ord.Internal_signature, &ord.Customer_id, &ord.Delivery_service,
			&ord.Shardkey, &ord.Sm_id, &ord.Date_created, &ord.Oof_shard,

			&del.Name, &del.Phone, &del.Zip, &del.City, &del.Address,
			&del.Region, &del.Email,

			&pay.Transaction, &pay.Request_id, &pay.Currency, &pay.Provider,
			&pay.Amount, &pay.Payment_dt, &pay.Bank, &pay.Delivery_cost,
			&pay.Goods_total, &pay.Custom_fee,

			&itm.Chrt_id, &itm.Price, &itm.Rid, &itm.Name, &itm.Sale,
			&itm.Size, &itm.Total_price, &itm.Nm_id, &itm.Brand, &itm.Status,
		); err != nil {
			return nil, err
		}

		if order.Order_uid == "" {
			order = ord
			order.Delivery = del
			order.Payment = pay
		}
		items = append(items, itm)
	}
	order.Items = items
	return &order, nil
}
