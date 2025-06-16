-- Order status table
CREATE TABLE IF NOT EXISTS
    order_status (
        id bigserial NOT NULL,
        name varchar(50) NOT NULL,
        description text NULL,
        CONSTRAINT order_status_pkey PRIMARY KEY (id),
        CONSTRAINT order_status_name_key UNIQUE (name)
    );

-- Shipping methods table
CREATE TABLE IF NOT EXISTS
    shipping_methods (
        id bigserial NOT NULL,
        name varchar(100) NOT NULL,
        description text NULL,
        price float8 NOT NULL,
        is_active boolean DEFAULT true NOT NULL,
        CONSTRAINT shipping_methods_pkey PRIMARY KEY (id),
        CONSTRAINT shipping_methods_name_key UNIQUE (name)
    );

-- Orders table
CREATE TABLE IF NOT EXISTS
    orders (
        id bigserial NOT NULL,
        user_id bigserial NOT NULL,
        order_number varchar(50) NOT NULL,
        status_id bigserial NOT NULL,
        shipping_method_id bigserial NOT NULL,
        shipping_addresses_id uuid NOT NULL,
        shipping_cost float8 NOT NULL,
        total_price float8 NOT NULL,
        final_price float8 NOT NULL,
        notes text NULL,
        created_at timestamptz (0) DEFAULT now () NOT NULL,
        updated_at timestamptz (0) DEFAULT now () NOT NULL,
        CONSTRAINT orders_pkey PRIMARY KEY (id),
        CONSTRAINT orders_order_number_key UNIQUE (order_number),
        CONSTRAINT orders_shipping_addresses_id_fkey FOREIGN KEY (shipping_addresses_id) REFERENCES shipping_addresses (id),
        CONSTRAINT orders_shipping_method_id_fkey FOREIGN KEY (shipping_method_id) REFERENCES shipping_methods (id),
        CONSTRAINT orders_status_id_fkey FOREIGN KEY (status_id) REFERENCES order_status (id),
        CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE INDEX idx_orders_user_id ON orders USING btree (user_id);

-- Order items table
CREATE TABLE IF NOT EXISTS
    order_items (
        id bigserial NOT NULL,
        order_id bigserial NOT NULL,
        product_id bigserial NOT NULL,
        toko_id bigserial NOT NULL,
        quantity int4 NOT NULL,
        price float8 NOT NULL,
        discount_price float8 NOT NULL,
        discount float8 NOT NULL,
        subtotal float8 NOT NULL,
        created_at timestamptz (0) DEFAULT now () NOT NULL,
        CONSTRAINT order_items_pkey PRIMARY KEY (id),
        CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
        CONSTRAINT order_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES products (id),
        CONSTRAINT order_items_toko_id_fkey FOREIGN KEY (toko_id) REFERENCES tokos (id)
    );

CREATE INDEX idx_order_items_order_id ON order_items USING btree (order_id);

CREATE INDEX idx_order_items_toko_id ON order_items USING btree (toko_id);

-- Order tracking table
CREATE TABLE IF NOT EXISTS
    order_tracking (
        id bigserial NOT NULL,
        order_id bigserial NOT NULL,
        status_id bigserial NOT NULL,
        notes text NULL,
        created_at timestamptz (0) DEFAULT now () NOT NULL,
        CONSTRAINT order_tracking_pkey PRIMARY KEY (id),
        CONSTRAINT order_tracking_order_id_fkey FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
        CONSTRAINT order_tracking_status_id_fkey FOREIGN KEY (status_id) REFERENCES order_status (id)
    );

CREATE INDEX idx_order_tracking_order_id ON order_tracking USING btree (order_id);

-- Initial data for order statuses
INSERT INTO 
    order_status (name, description)
VALUES
    (
        'Pending',
        'Order has been placed but not yet processed'
    ),
    (
        'Processing',
        'Order is being prepared for shipment'
    ),
    (
        'Shipped',
        'Order has been shipped to the customer'
    ),
    (
        'Delivered',
        'Order has been delivered to the customer'
    ),
    ('Cancelled', 'Order has been cancelled'),
    ('Refunded', 'Order has been refunded');

-- Initial data for shipping methods
INSERT INTO
    shipping_methods (name, description, price)
VALUES
    (
        'Standard Shipping',
        'Regular shipping with 3-5 business days delivery',
        9000
    ),
    (
        'Express Shipping',
        'Fast shipping with 1-2 business days delivery',
        12000
    ),
    (
        'Same Day Delivery',
        'Delivery on the same day for local orders',
        24000
    );