DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name='users' AND column_name='role_id'
    ) THEN
        ALTER TABLE users ADD COLUMN role_id INT REFERENCES roles (id) DEFAULT 1;
    END IF;
END$$;

UPDATE users
SET
    role_id = (
        SELECT id FROM roles WHERE name = 'user'
    )
WHERE role_id IS NULL;

ALTER TABLE users
ALTER COLUMN role_id DROP DEFAULT;

ALTER TABLE users
ALTER COLUMN role_id SET NOT NULL;