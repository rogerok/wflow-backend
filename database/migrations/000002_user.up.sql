ALTER TABLE users
    DROP COLUMN IF EXISTS age,
    DROP COLUMN IF EXISTS telegram_name;



ALTER TABLE users
    ADD COLUMN pseudonym_first_name varchar(50),
    ADD COLUMN pseudonym_last_name  varchar(50),
    ADD COLUMN social_telegram      varchar(255) unique,
    ADD COLUMN social_tiktok        varchar(255) unique,
    ADD COLUMN social_instagram     varchar(255) unique,
    ADD COLUMN social_vk            varchar(255) unique,
    ADD COLUMN born_date            date;


