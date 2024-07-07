create extension if not exists "uuid-ossp";

set timezone ="Europe/Moscow"


create table if not exists users
    (
        id uuid default uuid_generate_v4 () primary key
        email varchar(255) not null unique
        created_at timestamp with timezone default now ()
        updatedAt timestamp with timezone default now ()
        telegram_name varchar(255) unique
        first_name varchar(255) not null
    );

