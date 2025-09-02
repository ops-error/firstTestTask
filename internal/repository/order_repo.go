package repository

import (
	"context"
	"firstTestTask/internal/dto"

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
const queryInsertOrder = `
INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature,
                    customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (order_uid) DO UPDATE SET
`
const queryInsertDelivery = `
INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email,)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`
const queryInsertPayment = `
INSERT INTO payment (transaction, request_id, currency, provider, amount,
                     payment_dt, bank, delivery_cost, goods_total, custom_fee,)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`
const queryInsertItems = `
INSERT INTO items (order_uid, chrt_id, price, rid, name, sale,
                   size, total_price, nm_id, brand, status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
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
			//order
			&ord.Order_uid, &ord.Track_number, &ord.Entry, &ord.Locale,
			&ord.Internal_signature, &ord.Customer_id, &ord.Delivery_service,
			&ord.Shardkey, &ord.Sm_id, &ord.Date_created, &ord.Oof_shard,
			//delivery
			&del.Name, &del.Phone, &del.Zip, &del.City, &del.Address,
			&del.Region, &del.Email,
			//payment
			&pay.Transaction, &pay.Request_id, &pay.Currency, &pay.Provider,
			&pay.Amount, &pay.Payment_dt, &pay.Bank, &pay.Delivery_cost,
			&pay.Goods_total, &pay.Custom_fee,
			//item
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

func (repo *OrderRepo) SaveOrder(ctx context.Context, dto dto.OrderDTO) error {
	tx, err := repo.dataBase.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//order
	if _, err := tx.ExecContext(ctx, queryInsertOrder, dto.OrderUID, dto.TrackNumber,
		dto.Entry, dto.Locale, dto.InternalSignature, dto.CustomerID, dto.DeliveryService,
		dto.Shardkey, dto.SmID, dto.DateCreated, dto.OofShard); err != nil {
		return err
	}
	//delivery
	if _, err := tx.ExecContext(ctx, queryInsertDelivery, dto.OrderUID, dto.Delivery.Name,
		dto.Delivery.Phone, dto.Delivery.Zip, dto.Delivery.City, dto.Delivery.Address,
		dto.Delivery.Region, dto.Delivery.Email); err != nil {
		return err
	}
	//payment
	if _, err := tx.ExecContext(ctx, queryInsertPayment, dto.Payment.Transaction, dto.Payment.RequestID,
		dto.Payment.Currency, dto.Payment.Provider, dto.Payment.Amount, dto.Payment.PaymentDT,
		dto.Payment.Bank, dto.Payment.DeliveryCost, dto.Payment.GoodsTotal, dto.Payment.CustomFee); err != nil {
		return err
	}
	//items
	for _, i := range dto.Items {
		if _, err := tx.ExecContext(ctx, queryInsertItems, dto.OrderUID, i.ChrtID,
			i.Price, i.Rid, i.Name, i.Sale, i.Size, i.TotalPrice, i.NmID, i.Brand, i.Status); err != nil {
			return err
		}
	}
	return tx.Commit()
}
