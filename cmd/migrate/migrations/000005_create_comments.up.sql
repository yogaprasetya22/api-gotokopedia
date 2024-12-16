CREATE TABLE
    IF NOT EXISTS comments (
        id bigserial PRIMARY KEY,
        product_id bigserial NOT NULL,
        user_id bigserial NOT NULL,
        content text NOT NULL,
        created_at timestamp(0)
        with
            time zone NOT NULL DEFAULT NOW (),
            CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
            CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
    );