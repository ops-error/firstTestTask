-- Создание таблицы orders
CREATE TABLE orders (
                        order_uid VARCHAR(255) PRIMARY KEY,
                        track_number VARCHAR(255),
                        entry VARCHAR(50),
                        locale VARCHAR(10),
                        internal_signature VARCHAR(255),
                        customer_id VARCHAR(255),
                        delivery_service VARCHAR(255),
                        shardkey VARCHAR(10),
                        sm_id INTEGER,
                        date_created TIMESTAMP,
                        oof_shard VARCHAR(10)
);

-- Создание таблицы delivery
CREATE TABLE delivery (
                          id SERIAL PRIMARY KEY,
                          order_uid VARCHAR(255) REFERENCES orders(order_uid),
                          name VARCHAR(255),
                          phone VARCHAR(50),
                          zip VARCHAR(20),
                          city VARCHAR(255),
                          address TEXT,
                          region VARCHAR(255),
                          email VARCHAR(255)
);

-- Создание таблицы payment
CREATE TABLE payment (
                         transaction VARCHAR(255) PRIMARY KEY REFERENCES orders(order_uid),
                         request_id VARCHAR(255),
                         currency VARCHAR(10),
                         provider VARCHAR(100),
                         amount DECIMAL(10,2),
                         payment_dt BIGINT,
                         bank VARCHAR(100),
                         delivery_cost DECIMAL(10,2),
                         goods_total DECIMAL(10,2),
                         custom_fee DECIMAL(10,2)
);

-- Создание таблицы items
CREATE TABLE items (
                       id SERIAL PRIMARY KEY,
                       order_uid VARCHAR(255) REFERENCES orders(order_uid),
                       chrt_id BIGINT,
                       track_number VARCHAR(255),
                       price DECIMAL(10,2),
                       rid VARCHAR(255),
                       name VARCHAR(255),
                       sale INTEGER,
                       size VARCHAR(50),
                       total_price DECIMAL(10,2),
                       nm_id BIGINT,
                       brand VARCHAR(255),
                       status INTEGER
);

-- Вставка данных в таблицу orders
INSERT INTO orders (
    order_uid, track_number, entry, locale, internal_signature,
    customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
) VALUES (
             'b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 'en', '',
             'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1'
         );

-- Вставка данных в таблицу delivery
INSERT INTO delivery (
    order_uid, name, phone, zip, city, address, region, email
) VALUES (
             'b563feb7b2b84b6test', 'Test Testov', '+9720000000', '2639809',
             'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com'
         );

-- Вставка данных в таблицу payment
INSERT INTO payment (
    transaction, request_id, currency, provider, amount, payment_dt,
    bank, delivery_cost, goods_total, custom_fee
) VALUES (
             'b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727,
             'alpha', 1500, 317, 0
         );

-- Вставка данных в таблицу items
INSERT INTO items (
    order_uid, chrt_id, track_number, price, rid, name,
    sale, size, total_price, nm_id, brand, status
) VALUES (
             'b563feb7b2b84b6test', 9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest',
             'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202
         );