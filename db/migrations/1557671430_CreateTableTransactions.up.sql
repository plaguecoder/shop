create table transactions
(
  id          bigserial primary key    NOT NULL,
  customer_id INTEGER,
  date        TIMESTAMP with time zone NOT NULL default now(),
  amount      INTEGER                  NOT NULL,
  type        VARCHAR(256)             NOT NULL,
  description VARCHAR(256),
  created_at  TIMESTAMP with time zone NOT NULL default now()
);
