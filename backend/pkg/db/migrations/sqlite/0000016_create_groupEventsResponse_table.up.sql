CREATE TABLE IF NOT EXISTS groupEventResponses (
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		group_id INTEGER NOT NULL,
		response TEXT CHECK(response IN ('going', 'not going')),
		PRIMARY KEY (event_id, user_id, group_id),
		FOREIGN KEY (event_id) REFERENCES groupEvents(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
	);