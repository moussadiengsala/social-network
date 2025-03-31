-- UP SQL script
CREATE TABLE IF NOT EXISTS post (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			author_id INTEGER NOT NULL,
			image TEXT DEFAULT NULL,
			privacy TEXT CHECK(privacy IN ('private', 'public', 'almost_private')),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (author_id) REFERENCES user(id)
);