CREATE TABLE IF NOT EXISTS wallets
(
  id          SERIAL           PRIMARY KEY       ,
  "wallet_id" VARCHAR(255)     NOT NULL    UNIQUE,
  "balance"   DOUBLE PRECISION NOT NULL
);
