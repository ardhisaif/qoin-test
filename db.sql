CREATE TABLE customers(
    id VARCHAR( 50 ) PRIMARY KEY,
    address VARCHAR( 225 ),
    username VARCHAR ( 255 ) UNIQUE NOT NULL,
    email VARCHAR ( 255 ) UNIQUE NOT NULL,
    password VARCHAR ( 50 ) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE orders(
    id VARCHAR( 50 ) PRIMARY KEY,
    customer_id VARCHAR NOT NULL,
    order_number INT NOT NULL,
    is_payed BOOLEAN,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT fk_customer FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE order_details(
    id VARCHAR( 50 ) PRIMARY KEY,
    order_id VARCHAR NOT NULL,
    menu_id VARCHAR NOT NULL,
    total_item int,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT fk_menus FOREIGN KEY(menu_id) REFERENCES menus(id),
    CONSTRAINT fk_orders FOREIGN KEY(order_id) REFERENCES orders(id)
);

CREATE TABLE menus(
    id VARCHAR( 50 ) PRIMARY KEY,
    name varchar(255),
    price int,
    stock int,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

CREATE TABLE receipts(
    id VARCHAR( 50 ) PRIMARY KEY,
    customer_id VARCHAR NOT NULL,
    receipt_code VARCHAR(255) NOT NULL,
    items JSONB [],
    total_payment INT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT fk_customer FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE income_reports(
    id VARCHAR( 50 ) PRIMARY KEY,
    total_income int,
    total_customer int,
    total_sold_item int,
    total_order int,
    created_at TIMESTAMP NOT NULL
);
