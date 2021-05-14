create table if not exists entries
(
    id          serial primary key not null,
    token_id    integer            not null,
    document_id integer            not null,
    foreign key (token_id) references tokens (id) on delete cascade,
    foreign key (document_id) references documents (id) on delete cascade,
    unique (token_id, document_id)
);
