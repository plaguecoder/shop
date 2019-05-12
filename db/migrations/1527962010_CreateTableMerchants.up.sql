create table customers (
    id bigserial primary key not null,
    area varchar(256),
    name varchar(256) not null,
    phone varchar(20),
    description varchar(256),
    created_at timestamp with time zone not null default now()
);
