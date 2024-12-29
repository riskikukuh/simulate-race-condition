CREATE DATABASE simulation_race_condition;

CREATE TABLE users (
    id serial not null,
    name varchar(100) not null,
    email varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp default current_timestamp,
    primary key (id)
);

CREATE TABLE wallets (
    id serial not null,
    user_id int not null,
    balance bigint not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp default current_timestamp,
    primary key (id),
    CONSTRAINT user_id_unique UNIQUE (user_id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
