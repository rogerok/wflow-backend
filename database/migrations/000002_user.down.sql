alter table users
    drop column if exists born_date,
    drop column if exists pseudonym_first_name,
    drop column if exists pseudonym_last_name,
    drop column if exists social_tiktok,
    drop column if exists social_instagram,
    drop column if exists social_vk
;


alter table users
    rename column social_telegram to telegram_name;

alter table users
    add column if not exists age int;