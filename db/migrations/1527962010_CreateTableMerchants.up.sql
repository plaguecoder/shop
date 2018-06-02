create table merchants (
    id bigserial primary key not null,
    area varchar(256),
    name varchar(256) unique not null,
    phone varchar(20),
    created_at timestamp with time zone not null default now()
);
