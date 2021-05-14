create table if not exists tokens
(
    id   serial primary key  not null,
    text varchar(128) unique not null
);
