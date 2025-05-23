create extension if not exists "uuid-ossp";

set timezone = "Europe/Moscow";


DO
$$
    BEGIN
        CREATE TYPE role as ENUM ('admin', 'user');
    EXCEPTION
        WHEN duplicate_object THEN null;
    END
$$;


CREATE TABLE IF NOT EXISTS users
(
    id                   uuid         default uuid_generate_v4()                not null primary key,
    email                varchar(255)                                           not null unique,
    created_at           timestamp(0) default (now())::timestamp with time zone not null,
    updated_at           timestamp(0) default (now())::timestamp with time zone not null,
    first_name           varchar(255)                                           not null,
    last_name            varchar(255),
    middle_name          varchar(255),
    pseudonym_first_name varchar(50),
    pseudonym_last_name  varchar(50),
    social_telegram      varchar(255) unique,
    social_tiktok        varchar(255) unique,
    social_instagram     varchar(255) unique,
    social_vk            varchar(255) unique,
    born_date            timestamp(0) default NULL::timestamp without time zone,
    password             varchar(255)                                           not null,
    role                 role         default 'user'                            not null
);


CREATE TABLE IF NOT EXISTS sessions
(
    user_id       uuid unique                                            not null primary key,
    created_at    timestamp(0) default (now())::timestamp with time zone not null,
    updated_at    timestamp(0) default (now())::timestamp with time zone not null,
    refresh_token text                                                   not null,
    is_revoked    boolean                                                not null default false,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS books
(
    user_id     uuid                                                   not null,
    description varchar(255),
    id          uuid         default uuid_generate_v4()                not null primary key,
    created_at  timestamp(0) default (now())::timestamp with time zone not null,
    updated_at  timestamp(0) default (now())::timestamp with time zone not null,
    book_name   varchar(255)                                           not NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS goals
(
    book_id       uuid                                                   not null,
    created_at    timestamp(0) default (now())::timestamp with time zone not null,
    updated_at    timestamp(0) default (now())::timestamp with time zone not null,
    end_date      timestamp(0) default null::timestamp without time zone not null,
    goal_words    float8                                                   not null,
    id            uuid         default uuid_generate_v4()                not null primary key,
    is_finished   boolean      default false,
    start_date    timestamp(0) default null::timestamp without time zone not null,
    title         varchar(255)                                           not null,
    user_id       uuid                                                   not null,
    description   varchar(255),
    written_words float8                                                   not null,
    words_per_day float8                                                   not null,
    is_expired    boolean      default false,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE
);

CREATE INDEX idx_user_id ON goals (user_id);
CREATE INDEX idx_book_id ON goals (book_id);


CREATE TABLE IF NOT EXISTS reports
(
    book_id      uuid                                                   not null,
    goal_id      uuid                                                   not null,
    created_at   timestamp(0) default (now())::timestamp with time zone not null,
    updated_at   timestamp(0) default (now())::timestamp with time zone not null,
    words_amount float8                                                   not null,
    id           uuid         default uuid_generate_v4()                not null primary key,
    title        varchar(255)                                           not null,
    user_id      uuid                                                   not null,
    description  varchar(255),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    CONSTRAINT fk_goal FOREIGN KEY (goal_id) REFERENCES goals (id) ON DELETE CASCADE
);

CREATE INDEX idx_books_user_id ON books (user_id);
CREATE INDEX idx_goals_book_id ON goals (book_id);
CREATE INDEX idx_reports_book_id ON reports (book_id);
CREATE INDEX idx_reports_goal_id ON reports (goal_id);
CREATE INDEX idx_reports_user_id ON reports (user_id);


CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER update_users_timestamp
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_sessions_timestamp
    BEFORE UPDATE
    ON sessions
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_books_timestamp
    BEFORE UPDATE
    ON books
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_goals_timestamp
    BEFORE UPDATE
    ON goals
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_reports_timestamp
    BEFORE UPDATE
    ON reports
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();