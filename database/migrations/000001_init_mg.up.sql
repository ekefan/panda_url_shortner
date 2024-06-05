-- Up Migration: Create the users table
-- Create the users table
CREATE TABLE users(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add index on user name
CREATE INDEX idx_user_name ON users(name);

-- Create the urls table
CREATE TABLE urls(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    owner INTEGER NOT NULL,
    short_code TEXT NOT NULL UNIQUE,
    long_url TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00',
    FOREIGN KEY (owner) REFERENCES users(id),
    UNIQUE (owner, short_code),
    UNIQUE (owner, long_url)
);

-- Add index on url owner
CREATE INDEX idx_url_owner ON urls(owner);

