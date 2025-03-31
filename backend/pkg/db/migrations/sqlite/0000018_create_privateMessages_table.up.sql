CREATE TABLE IF NOT EXISTS privateMessage (
		id INTEGER PRIMARY KEY,
		sender_id INT,
		receiver_id INT,
		content TEXT,
		status TEXT CHECK (status IN ('read', 'unread')) DEFAULT 'unread',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (sender_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (receiver_id) REFERENCES user(id) ON DELETE CASCADE
	);