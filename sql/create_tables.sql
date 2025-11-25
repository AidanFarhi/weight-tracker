DROP TABLE IF EXISTS user, weight_entry, session CASCADE;

CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE,
    password TEXT
);

CREATE TABLE weight_entry (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES user(id),
    weight INT,
    entry_date TEXT,
    category TEXT
);

CREATE TABLE session (
    id TEXT UNIQUE,
    user_id INT REFERENCES user(id)
);