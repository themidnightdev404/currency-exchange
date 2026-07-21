CREATE TABLE IF NOT EXISTS currencies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code TEXT NOT NULL UNIQUE,
    fullname TEXT NOT NULL,
    sign TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS exchange_rates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    base_currency_id INTEGER NOT NULL,
    target_currency_id INTEGER NOT NULL,
    rate DECIMAL(18, 6) NOT NULL,
    FOREIGN KEY (base_currency_id) REFERENCES currencies (id),
    FOREIGN KEY (target_currency_id) REFERENCES currencies (id),
    UNIQUE (base_currency_id, target_currency_id)
);

INSERT INTO
    currencies (code, fullname, sign)
VALUES
    ('USD', 'United States Dollar', '$');

INSERT INTO
    currencies (code, fullname, sign)
VALUES
    ('EUR', 'Euro', '€');

INSERT INTO
    currencies (code, fullname, sign)
VALUES
    ('RUB', 'Russian Ruble', '₽');