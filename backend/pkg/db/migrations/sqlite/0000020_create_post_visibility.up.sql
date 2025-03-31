CREATE TABLE IF NOT EXISTS post_visibility (
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (post_id) REFERENCES post(id),
		FOREIGN KEY (user_id) REFERENCES user(id)
);