CREATE TABLE wallets
(
  id               SERIAL           PRIMARY KEY,
  "wallet_id"      VARCHAR(255)     NOT NULL,
  "operation_type" VARCHAR(255)     NOT NULL,
  "amount"         DOUBLE PRECISION NOT NULL
);
