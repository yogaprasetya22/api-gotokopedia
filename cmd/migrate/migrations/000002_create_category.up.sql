CREATE TABLE
    IF NOT EXISTS category (
        id bigserial PRIMARY KEY,
        name TEXT NOT NULL,
        slug TEXT NOT NULL,
        description TEXT
    );