CREATE TABLE IF NOT EXISTS notifications (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender_id INT,
		receiver_id INT,
		notification_type TEXT CHECK(notification_type IN ('follow_request', 'message', 'groups_invited', 'groups_requested', 'groups_events')),
		group_id INT DEFAULT 0,
		event_id INT DEFAULT 0,
		status TEXT CHECK(status IN ('read', 'unread')) DEFAULT 'unread',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (receiver_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (sender_id) REFERENCES user(id) ON DELETE CASCADE
	);