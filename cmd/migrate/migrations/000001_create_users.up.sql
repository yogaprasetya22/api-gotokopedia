CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE
    IF NOT EXISTS users (
        id bigserial PRIMARY KEY,
        google_id VARCHAR(255),
        username VARCHAR(255) NOT NULL,
        email citext UNIQUE NOT NULL,
        password bytea,
        is_active BOOLEAN NOT NULL DEFAULT FALSE,
        picture VARCHAR(255) NOT NULL DEFAULT 'https://upload.wikimedia.org/wikipedia/commons/9/99/Sample_User_Icon.png',
        created_at timestamp(0)
        with
            time zone NOT NULL DEFAULT now ()
    );