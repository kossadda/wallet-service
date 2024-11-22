CREATE TABLE wallets
(
  id       SERIAL           PRIMARY KEY,
  "uuid"   UUID             NOT NULL,
  "type"   VARCHAR(255)     NOT NULL,
  "amount" DOUBLE PRECISION NOT NULL
);
