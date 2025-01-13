CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS books
(
    user_id     uuid         NOT NULL,
    description VARCHAR(255),
    id          uuid PRIMARY KEY         DEFAULT uuid_generate_v4(),
    created_at  timestamp with time zone DEFAULT now(),
    updated_at  timestamp with time zone DEFAULT now(),
    book_name   VARCHAR(255) NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);
