CREATE TABLE crypto_prices (
                               symbolFrom VARCHAR(255) NOT NULL,
                               symbolTo VARCHAR(255) NOT NULL,
                               price NUMERIC(20, 10) NOT NULL,

                               CONSTRAINT crypto_prices_pkey PRIMARY KEY (symbolFrom, symbolTo)
);