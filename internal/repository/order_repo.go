package repository

import (
	"context"
	"firstTestTask/internal/domain"
	"log"

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
INSERT INTO orders (
    order_uid, track_number, entry, locale, internal_signature,
    customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
ON CONFLICT (order_uid) DO UPDATE SET
    track_number   = EXCLUDED.track_number,
    entry          = EXCLUDED.entry,
    locale         = EXCLUDED.locale,
    internal_signature = EXCLUDED.internal_signature,
    customer_id    = EXCLUDED.customer_id,
    delivery_service = EXCLUDED.delivery_service,
    shardkey       = EXCLUDED.shardkey,
    sm_id          = EXCLUDED.sm_id,
    date_created   = EXCLUDED.date_created,
    oof_shard      = EXCLUDED.oof_shard;
`
const queryInsertDelivery = `
INSERT INTO delivery (
    order_uid, name, phone, zip, city, address, region, email
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`
const queryInsertPayment = `
INSERT INTO payment (
    transaction, request_id, currency, provider, amount,
    payment_dt, bank, delivery_cost, goods_total, custom_fee
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (transaction) DO NOTHING;
`
const queryInsertItems = `
INSERT INTO items (
    order_uid, chrt_id, price, rid, name, sale,
    size, total_price, nm_id, brand, status
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
`

func (repo *OrderRepo) GetFullOrder(ctx context.Context, uid string) (*domain.OrderDTO, error) {
	rows, err := repo.dataBase.QueryxContext(ctx, queryFullOrder, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		order domain.OrderDTO
		items []domain.ItemDTO
	)
	for rows.Next() {
		var (
			ord domain.OrderDTO
			del domain.DeliveryDTO
			pay domain.PaymentDTO
			itm domain.ItemDTO
		)

		if err := rows.Scan(
			//order
			&ord.OrderUID, &ord.TrackNumber, &ord.Entry, &ord.Locale,
			&ord.InternalSignature, &ord.CustomerID, &ord.DeliveryService,
			&ord.Shardkey, &ord.SmID, &ord.DateCreated, &ord.OofShard,
			//delivery
			&del.Name, &del.Phone, &del.Zip, &del.City, &del.Address,
			&del.Region, &del.Email,
			//payment
			&pay.Transaction, &pay.RequestID, &pay.Currency, &pay.Provider,
			&pay.Amount, &pay.PaymentDT, &pay.Bank, &pay.DeliveryCost,
			&pay.GoodsTotal, &pay.CustomFee,
			//item
			&itm.ChrtID, &itm.Price, &itm.Rid, &itm.Name, &itm.Sale,
			&itm.Size, &itm.TotalPrice, &itm.NmID, &itm.Brand, &itm.Status,
		); err != nil {
			return nil, err
		}

		if order.OrderUID == "" {
			order = ord
			order.Delivery = del
			order.Payment = pay
		}
		items = append(items, itm)
	}
	order.Items = items
	return &order, nil
}

func (repo *OrderRepo) SaveOrder(ctx context.Context, dto domain.OrderDTO) error {
	tx, err := repo.dataBase.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//order
	log.Println("order")
	if _, err := tx.ExecContext(ctx, queryInsertOrder, dto.OrderUID, dto.TrackNumber,
		dto.Entry, dto.Locale, dto.InternalSignature, dto.CustomerID, dto.DeliveryService,
		dto.Shardkey, dto.SmID, dto.DateCreated, dto.OofShard); err != nil {
		return err
	}
	//delivery
	log.Println("delivery")
	if _, err := tx.ExecContext(ctx, queryInsertDelivery, dto.OrderUID, dto.Delivery.Name,
		dto.Delivery.Phone, dto.Delivery.Zip, dto.Delivery.City, dto.Delivery.Address,
		dto.Delivery.Region, dto.Delivery.Email); err != nil {
		return err
	}
	//payment
	log.Println("payment")
	if _, err := tx.ExecContext(ctx, queryInsertPayment, dto.Payment.Transaction, dto.Payment.RequestID,
		dto.Payment.Currency, dto.Payment.Provider, dto.Payment.Amount, dto.Payment.PaymentDT,
		dto.Payment.Bank, dto.Payment.DeliveryCost, dto.Payment.GoodsTotal, dto.Payment.CustomFee); err != nil {
		return err
	}
	//items
	log.Println("items")
	for _, i := range dto.Items {
		if _, err := tx.ExecContext(ctx, queryInsertItems, dto.OrderUID, i.ChrtID,
			i.Price, i.Rid, i.Name, i.Sale, i.Size, i.TotalPrice, i.NmID, i.Brand, i.Status); err != nil {
			return err
		}
	}
	return tx.Commit()
}
