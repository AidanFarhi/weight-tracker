DROP TABLE IF EXISTS user_account, weight_entry, user_session CASCADE;

CREATE TABLE user_account (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE,
    password TEXT
);

CREATE TABLE weight_entry (
    id SERIAL PRIMARY KEY,
    user_account_id INT REFERENCES user_account(id),
    weight INT,
    entry_date TEXT,
    category TEXT
);

CREATE TABLE user_session (
    id TEXT UNIQUE,
    user_account_id INT REFERENCES user_account(id)
);