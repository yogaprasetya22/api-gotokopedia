CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    IF NOT EXISTS carts (
        id bigserial NOT NULL,
        user_id bigserial NOT NULL,
        created_at timestamptz (0) DEFAULT now () NOT NULL,
        updated_at timestamptz (0) DEFAULT now () NOT NULL,
        CONSTRAINT carts_pkey PRIMARY KEY (id),
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE INDEX idx_cart_user_id ON carts USING btree (user_id);

CREATE TABLE
    IF NOT EXISTS cart_stores (
        id uuid DEFAULT uuid_generate_v4 () NOT NULL,
        cart_id bigint NOT NULL,
        toko_id bigint NOT NULL,
        created_at timestamptz (0) DEFAULT now () NOT NULL,
        CONSTRAINT cart_stores_pkey PRIMARY KEY (id),
        CONSTRAINT fk_cart FOREIGN KEY (cart_id) REFERENCES carts (id) ON DELETE CASCADE,
        CONSTRAINT fk_toko FOREIGN KEY (toko_id) REFERENCES tokos (id)
    );

CREATE INDEX idx_cart_stores_cart_toko ON cart_stores USING btree (cart_id, toko_id);

CREATE TABLE
    IF NOT EXISTS cart_items (
        id uuid DEFAULT uuid_generate_v4 () NOT NULL,
        cart_id bigserial NOT NULL,
        cart_store_id uuid NOT NULL,
        product_id bigserial NOT NULL,
        quantity int4 DEFAULT 1 NOT NULL,
        created_at timestamptz (0) DEFAULT now () NOT NULL,
        updated_at timestamptz (0) DEFAULT now () NOT NULL,
        CONSTRAINT cart_items_pkey PRIMARY KEY (id),
        CONSTRAINT fk_cart FOREIGN KEY (cart_id) REFERENCES carts (id) ON DELETE CASCADE,
        CONSTRAINT fk_cart_store FOREIGN KEY (cart_store_id) REFERENCES cart_stores (id) ON DELETE CASCADE,
        CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
    );

CREATE INDEX idx_cart_items_cart_product ON cart_items USING btree (cart_id, product_id);

CREATE INDEX idx_cart_items_cart_store ON cart_items USING btree (cart_store_id);