-- Payment methods table
CREATE TABLE
    payment_methods (
        id bigserial NOT NULL,
        name varchar(100) NOT NULL,
        description text NULL,
        midtrans_code varchar(50) NULL,
        is_active boolean DEFAULT true NOT NULL,
        CONSTRAINT payment_methods_pkey PRIMARY KEY (id),
        CONSTRAINT payment_methods_name_key UNIQUE (name)
    );

-- Payments table
CREATE TABLE
    payments (
        id bigserial NOT NULL,
        order_id bigserial NOT NULL,
        amount float8 NOT NULL,
        payment_method_id bigserial NOT NULL,
        transaction_id varchar(255) NULL,
        status varchar(50) NOT NULL,
        payment_date timestamptz (0) NULL,
        created_at timestamptz (0) DEFAULT now () NOT NULL,
        updated_at timestamptz (0) DEFAULT now () NOT NULL,
        CONSTRAINT payments_pkey PRIMARY KEY (id),
        CONSTRAINT payments_order_id_fkey FOREIGN KEY (order_id) REFERENCES orders (id),
        CONSTRAINT payments_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id)
    );

CREATE INDEX idx_payments_order_id ON payments USING btree (order_id);

-- Add payment method column to orders table
ALTER TABLE orders
ADD COLUMN payment_method_id bigserial;

ALTER TABLE orders ADD CONSTRAINT orders_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id);

-- Initial data for payment methods
INSERT INTO
    payment_methods (name, description, midtrans_code, is_active)
VALUES
    (
        'Bank Transfer BCA',
        'Transfer via BCA Virtual Account',
        'bank_transfer.bca',
        true
    ),
    (
        'Bank Transfer BNI',
        'Transfer via BNI Virtual Account',
        'bank_transfer.bni',
        true
    ),
    (
        'Bank Transfer BRI',
        'Transfer via BRI Virtual Account',
        'bank_transfer.bri',
        true
    ),
    ('Gopay', 'Pembayaran via Gopay', 'gopay', true),
    ('QRIS', 'Pembayaran via QRIS', 'qris', true),
    (
        'Alfamart',
        'Pembayaran via Alfamart',
        'cstore.alfamart',
        true
    ),
    (
        'Indomaret',
        'Pembayaran via Indomaret',
        'cstore.indomaret',
        true
    );