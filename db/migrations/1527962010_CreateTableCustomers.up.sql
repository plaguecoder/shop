create table areas
(
  id         bigserial primary key    NOT NULL,
  name       VARCHAR(256) UNIQUE      NOT NULL,
  created_at TIMESTAMP with time zone NOT NULL default now()
);

create table customers
(
  id          bigserial primary key        NOT NULL,
  area_id     BIGINT REFERENCES areas (id) NOT NULL,
  name        VARCHAR(256)                 NOT NULL,
  phone       VARCHAR(20),
  description VARCHAR(256),
  created_at  TIMESTAMP with time zone     NOT NULL default now(),
  CONSTRAINT unique_name_in_area UNIQUE (name, area_id)
);

create table transactions
(
  id          bigserial primary key            NOT NULL,
  customer_id BIGINT REFERENCES customers (id) NOT NULL,
  date        TIMESTAMP with time zone         NOT NULL default now(),
  amount      INTEGER                          NOT NULL,
  type        VARCHAR(256)                     NOT NULL,
  description VARCHAR(256),
  created_at  TIMESTAMP with time zone         NOT NULL default now()
);