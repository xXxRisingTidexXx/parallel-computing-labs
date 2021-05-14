create table if not exists users
(
    id       serial primary key not null,
    username varchar(64) unique not null,
    password varchar(256)       not null,
    role     varchar(20)        not null
);
