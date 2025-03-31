CREATE TABLE IF NOT EXISTS groupMembers (
		group_id INTEGER,
		user_id INTEGER,
		role TEXT CHECK(role IN ('admin', 'user')),
		status TEXT CHECK(status IN ('invited', 'accepted', 'requested', 'rejected')),
		joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (group_id, user_id),
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
	);