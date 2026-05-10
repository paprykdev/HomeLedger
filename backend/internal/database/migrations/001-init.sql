PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME
);

CREATE TABLE IF NOT EXISTS accounts (
    id TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,

    currency TEXT NOT NULL CHECK (
        currency IN ('PLN', 'USD', 'EUR')
    ),

    user_id TEXT NOT NULL,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,

    FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id TEXT PRIMARY KEY NOT NULL,

    type TEXT NOT NULL CHECK (
        type IN ('income', 'expense')
    ),

    amount REAL NOT NULL CHECK (
        amount > 0
    ),

    description TEXT,

    category TEXT NOT NULL,

    account_id TEXT NOT NULL,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,

    FOREIGN KEY(account_id) REFERENCES accounts(id)
);

CREATE TABLE IF NOT EXISTS scheduled_transactions (
    id TEXT PRIMARY KEY NOT NULL,

    type TEXT NOT NULL CHECK (
        type IN ('income', 'expense')
    ),

    amount REAL NOT NULL CHECK (
        amount > 0
    ),

    frequency TEXT NOT NULL CHECK (
        frequency IN (
            'once',
            'daily',
            'weekly',
            'monthly',
            'yearly'
        )
    ),

    next_run_at DATETIME NOT NULL,

    description TEXT,

    category TEXT NOT NULL,

    account_id TEXT NOT NULL,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,

    FOREIGN KEY(account_id) REFERENCES accounts(id)
);
