CREATE TABLE IF NOT EXISTS groupMessageStatus (
		user_id INT,
		group_message_id INT,
		status TEXT CHECK (status IN ('read', 'unread')) DEFAULT 'unread',
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (group_message_id) REFERENCES groupMessage(id) ON DELETE CASCADE
	);