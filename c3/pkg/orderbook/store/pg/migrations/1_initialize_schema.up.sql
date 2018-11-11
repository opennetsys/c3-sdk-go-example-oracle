CREATE TABLE IF NOT EXISTS public.accounts
(
    id bigserial NOT NULL,
    address text NOT NULL,
    chain text NOT NULL,
    UNIQUE(address, chain),
    CONSTRAINT accounts_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS account_address_index ON accounts(address);



CREATE TABLE IF NOT EXISTS public.balances
(
    account_id bigint NOT NULL REFERENCES accounts (id),
    currency text COLLATE pg_catalog."default" NOT NULL,
    balance numeric NOT NULL,
    CONSTRAINT balances_pkey PRIMARY KEY (account_id, currency)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;



DO $$ BEGIN
    CREATE TYPE order_t AS ENUM('BID', 'ASK');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;



CREATE TABLE IF NOT EXISTS public.orderbook
(
    id bigserial NOT NULL,
    account_id bigint NOT NULL REFERENCES accounts (id),
    symbol text COLLATE pg_catalog."default" NOT NULL,
    type order_t NOT NULL,
    rate numeric NOT NULL,
    quantity numeric NOT NULL,
    CONSTRAINT orderbook_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS account_orderbook_index ON orderbook(account_id, symbol, type);
CREATE INDEX IF NOT EXISTS symbol_orderbook_index ON orderbook(symbol, type, rate);
