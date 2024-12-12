CREATE TABLE
    IF NOT EXISTS "product" (
        id bigserial PRIMARY KEY,
        name TEXT NOT NULL,
        slug TEXT NOT NULL,
        description TEXT,
        price DOUBLE PRECISION NOT NULL,
        discount_price DOUBLE PRECISION NOT NULL,
        discount DOUBLE PRECISION NOT NULL,
        rating DOUBLE PRECISION NOT NULL,
        estimation TEXT NOT NULL,
        stock INTEGER NOT NULL,
        sold INTEGER NOT NULL,
        is_for_sale BOOLEAN NOT NULL,
        is_approved BOOLEAN NOT NULL,
        image_urls TEXT[],
        category_id bigint NOT NULL,
        toko_id bigint NOT NULL,
        created_at timestamp(0)
        with
            time zone NOT NULL DEFAULT NOW (),
        updated_at timestamp(0)
        with
            time zone NOT NULL DEFAULT NOW (),
            CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES "category" (id),
            CONSTRAINT fk_toko FOREIGN KEY (toko_id) REFERENCES "toko" (id)
    );