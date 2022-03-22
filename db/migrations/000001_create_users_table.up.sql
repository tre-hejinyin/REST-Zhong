ALTER DATABASE template SET timezone TO 'Asia/Tokyo';

create table if not exists users
(
    id              INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name      TEXT NOT NULL CHECK(length(first_name) >= 1 AND length(first_name) <= 64),
    last_name       TEXT NOT NULL CHECK(length(last_name) >= 1 AND length(last_name) <= 64),
    email           TEXT NOT NULL UNIQUE,
    password        VARCHAR(100) NOT NULL ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);
