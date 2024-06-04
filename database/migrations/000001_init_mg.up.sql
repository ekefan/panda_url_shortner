-- Up Migration: Create the users table
CREATE TABLE users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Up Migration: Create the urls table
CREATE TABLE urls(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    owner INTEGER NOT NULL,
    short_code TEXT NOT NULL,
    long_url TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT 0,
    FOREIGN KEY (owner) REFERENCES users(id),
    UNIQUE (owner, short_code),
    UNIQUE (owner, long_url)
);
