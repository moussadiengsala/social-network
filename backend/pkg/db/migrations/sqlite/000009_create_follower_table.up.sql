CREATE TABLE IF NOT EXISTS follower (
		id INTEGER PRIMARY KEY,
		follower_id INTEGER,
		following_id INTEGER,
		status TEXT CHECK (status IN ('pending', 'accept', 'reject')) DEFAULT 'pending',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (follower_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (following_id) REFERENCES user(id) ON DELETE CASCADE,
		CONSTRAINT unique_follow_pair UNIQUE (follower_id, following_id),
		CONSTRAINT different_ids CHECK (follower_id != following_id)
	);