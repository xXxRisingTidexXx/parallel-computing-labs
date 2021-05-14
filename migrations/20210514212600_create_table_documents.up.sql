create table if not exists documents
(
    id      serial primary key not null,
    text    varchar(1024)      not null,
    site_id integer            not null,
    foreign key (site_id) references sites (id) on delete cascade
);
