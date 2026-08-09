CREATE TABLE IF NOT EXISTS currencies (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Code TEXT NOT NULL UNIQUE,
    Name TEXT NOT NULL,
    Sign TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS ExchangeRates (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    BaseCurrencyId INTEGER NOT NULL,
    TargetCurrencyId INTEGER NOT NULL,
    Rate REAL NOT NULL,
    FOREIGN KEY (BaseCurrencyId) REFERENCES currencies (ID),
    FOREIGN KEY (TargetCurrencyId) REFERENCES currencies (ID),
    UNIQUE (BaseCurrencyId, TargetCurrencyId)
);

CREATE INDEX IF NOT EXISTS idx_base_currency ON ExchangeRates(BaseCurrencyId);