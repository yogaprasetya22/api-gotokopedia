CREATE TABLE
    IF NOT EXISTS "toko" (
        id bigserial PRIMARY KEY,
        user_id bigserial NOT NULL,
        slug TEXT NOT NULL,
        name TEXT NOT NULL,
        image_profile TEXT,
        country TEXT NOT NULL,
        created_at timestamp(0)
        with
            time zone NOT NULL DEFAULT NOW (),
            CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user" (id)
    );