CREATE TABLE users
(
    id   SERIAL PRIMARY KEY,
    username varchar(255),
    password bytea

);
CREATE TABLE task
(
    id   SERIAL PRIMARY KEY,
    title varchar(255),
    comment varchar(255),
    done bool,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);