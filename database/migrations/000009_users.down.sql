ALTER TABLE users
    ALTER COLUMN created_at TYPE timestamp with time zone USING created_at::timestamp with time zone;

ALTER TABLE users
    ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE users
    ALTER COLUMN created_at SET NOT NULL;