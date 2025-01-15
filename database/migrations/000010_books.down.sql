ALTER TABLE books
    ALTER COLUMN created_at TYPE timestamp with time zone USING created_at::timestamp with time zone;

ALTER TABLE books
    ALTER COLUMN created_at SET DEFAULT now();

ALTER TABLE books
    ALTER COLUMN created_at SET NOT NULL;


ALTER TABLE books
    ALTER COLUMN updated_at TYPE timestamp with time zone USING created_at::timestamp with time zone;

ALTER TABLE books
    ALTER COLUMN updated_at SET DEFAULT now();

ALTER TABLE books
    ALTER COLUMN updated_at SET NOT NULL;