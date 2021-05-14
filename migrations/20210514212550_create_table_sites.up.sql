create table if not exists sites
(
    id     serial primary key  not null,
    domain varchar(256) unique not null
);
