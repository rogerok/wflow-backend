ALTER TABLE users
    ADD COLUMN IF NOT EXISTS password varchar(255);