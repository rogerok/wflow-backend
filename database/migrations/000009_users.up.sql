ALTER TABLE users
    ALTER COLUMN created_at TYPE timestamp(0) with time zone USING created_at::timestamp(0) with time zone;

ALTER TABLE users
    ALTER COLUMN created_at SET DEFAULT (now())::timestamp(0) with time zone;

ALTER TABLE users
    ALTER COLUMN created_at SET NOT NULL;