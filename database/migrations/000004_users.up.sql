ALTER TABLE users
    ALTER COLUMN password  TYPE varchar(255),
    ALTER COLUMN born_date TYPE timestamp;

ALTER TABLE users
    ALTER COLUMN password SET NOT NULL;