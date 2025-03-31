-- Up SQL script
CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE,
    username TEXT UNIQUE,
    date_of_birth DATE,
    gender TEXT CHECK(gender IN ('male', 'female')),
    bio TEXT,
    avatar TEXT,
    password TEXT,
    privacy TEXT CHECK(privacy IN ('public', 'private')) DEFAULT 'public',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);