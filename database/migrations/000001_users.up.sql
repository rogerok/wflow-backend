create extension if not exists "uuid-ossp";

set timezone ="Europe/Moscow";


create table if not exists users
    (
        id uuid default uuid_generate_v4() not null primary key,
        email varchar(255) not null unique,
        created_at timestamp with time zone default now(),
        updated_at timestamp with time zone default now(),
        telegram_name varchar(255) unique,
        first_name varchar(255) not null,
        last_name varchar(255),
        middle_name varchar(255),
        age int
    );
