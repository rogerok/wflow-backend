create extension if not exists "uuid-ossp";

set timezone = "Europe/Moscow";


CREATE TYPE role AS ENUM ('admin', 'user');

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
    refresh_token TEXT                                                   not null,
    is_revoked    BOOLEAN                                                not null default false,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE IF NOT EXISTS books
(
    user_id     uuid                                                   not null,
    description varchar(255),
    id          uuid         default uuid_generate_v4()                not null primary key,
    created_at  timestamp(0) default (now())::timestamp with time zone not null,
    updated_at  timestamp(0) default (now())::timestamp with time zone not null,
    book_name   VARCHAR(255)                                           not NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE IF NOT EXISTS goals
(
    book_id       uuid                                                   not null,
    created_at    timestamp(0) default (now())::timestamp with time zone not null,
    updated_at    timestamp(0) default (now())::timestamp with time zone not null,
    end_date      timestamp(0) default null::timestamp without time zone not null,
    goal_words    real                                                   not null,
    id            uuid         default uuid_generate_v4()                not null primary key,
    is_finished   boolean      default false,
    start_date    timestamp(0) default null::timestamp without time zone not null,
    title         varchar(255)                                           not null,
    user_id       uuid                                                   not null,
    description   varchar(255),
    written_words real                                                   not null,
    words_per_day real                                                   not null,
    is_expired    boolean      default false,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id)
);

CREATE INDEX idx_user_id ON goals (user_id);
CREATE INDEX idx_book_id ON goals (book_id);


CREATE TABLE IF NOT EXISTS reports
(
    book_id      uuid                                                   not null,
    goal_id      uuid                                                   not null,
    created_at   timestamp(0) default (now())::timestamp with time zone not null,
    updated_at   timestamp(0) default (now())::timestamp with time zone not null,
    words_amount real                                                   not null,
    id           uuid         default uuid_generate_v4()                not null primary key,
    title        varchar(255)                                           not null,
    user_id      uuid                                                   not null,
    description  varchar(255),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books (id),
    CONSTRAINT fk_goal FOREIGN KEY (goal_id) REFERENCES goals (id)
);


CREATE INDEX idx_reports_book_id ON reports (book_id);
CREATE INDEX idx_reports_goal_id ON reports (goal_id);
CREATE INDEX idx_reports_user_id ON reports (user_id);