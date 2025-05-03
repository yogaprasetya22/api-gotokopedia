CREATE TABLE
    shipping_addresses (
        id uuid DEFAULT uuid_generate_v4 () NOT NULL,
        user_id bigserial NOT NULL,
        label varchar(100) NOT NULL,
        recipient_name varchar(100) NOT NULL,
        recipient_phone varchar(50) NOT NULL,
        address_line1 text NOT NULL,
        note_for_courier text NULL,
        created_at timestamptz (0) DEFAULT now () NOT NULL,
        updated_at timestamptz (0) DEFAULT now () NOT NULL,
        CONSTRAINT shipping_address_pkey PRIMARY KEY (id),
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );