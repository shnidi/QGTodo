CREATE TABLE users
(
    id   SERIAL PRIMARY KEY,
    username varchar(255),
    password bytea,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
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

CREATE TABLE groups(
    id SERIAL PRIMARY KEY,
    title varchar(255),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE usergroup(
        id SERIAL PRIMARY KEY,
        fk_user  SERIAL REFERENCES users(id) ON DELETE RESTRICT,
        fk_group  SERIAL REFERENCES groups(id) ON DELETE RESTRICT
);