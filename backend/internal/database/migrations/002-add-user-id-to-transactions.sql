CREATE TABLE transactions_new (
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
    user_id TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    FOREIGN KEY(account_id) REFERENCES accounts(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);

INSERT INTO transactions_new (
    id,
    type,
    amount,
    description,
    category,
    account_id,
    user_id,
    created_at,
    updated_at,
    deleted_at
)
SELECT
    t.id,
    t.type,
    t.amount,
    t.description,
    t.category,
    t.account_id,
    a.user_id,
    t.created_at,
    t.updated_at,
    t.deleted_at
FROM transactions t
JOIN accounts a ON a.id = t.account_id;

DROP TABLE transactions;
ALTER TABLE transactions_new RENAME TO transactions;

CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);
