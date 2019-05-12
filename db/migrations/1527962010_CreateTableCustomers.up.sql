create table customers
(
  id          bigserial primary key    NOT NULL,
  area        VARCHAR(256),
  name        VARCHAR(256)             NOT NULL,
  phone       VARCHAR(20),
  description VARCHAR(256),
  created_at  TIMESTAMP with time zone NOT NULL default now()
);
