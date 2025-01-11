
CREATE TABLE IF NOT EXISTS sessions (
    user_id uuid UNIQUE NOT NULL,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now(),
    expires_at timestamp with time zone,
    refresh_token TEXT NOT NULL,
    is_revoked BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
)