-- Migration for shipping_addresses table
CREATE TABLE
    IF NOT EXISTS shipping_addresses (
        id uuid DEFAULT uuid_generate_v4 () NOT NULL,
        user_id bigserial NOT NULL,
        is_default boolean DEFAULT false NOT NULL,
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

-- Add unique partial index for default address
CREATE UNIQUE INDEX idx_one_default_address_per_user ON shipping_addresses (user_id)
WHERE
    is_default = true;

-- Add foreign key for default_shipping_address_id after shipping_addresses is created
ALTER TABLE users ADD CONSTRAINT fk_default_shipping_address FOREIGN KEY (default_shipping_address_id) REFERENCES shipping_addresses (id) ON DELETE SET NULL;