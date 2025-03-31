package internals

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/*
	TablesCreation initializes the database schema by creating necessary tables.
	It takes an instance of *sql.DB as a parameter and ensures that the required
	tables are created with their specified structures.
*/
// ...

func TablesCreation(instanceOfDb *sql.DB) {
	var err error
	if instanceOfDb == nil {
		log.Println("Unable to reach the database")
		return
	}

	// Create the "post" table
	_, err = instanceOfDb.Exec(`
		CREATE TABLE IF NOT EXISTS post (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			author_id INTEGER NOT NULL,
			image TEXT DEFAULT NULL,
			privacy TEXT CHECK(privacy IN ('private', 'public', 'almost_private')),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (author_id) REFERENCES user(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Post visivility for almost private posts
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS post_visibility (
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (post_id) REFERENCES post(id),
		FOREIGN KEY (user_id) REFERENCES user(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "user" table
	_, err = instanceOfDb.Exec(`
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
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "comments" table
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		entrie_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		image TEXT DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (entrie_id) REFERENCES post(id),
		FOREIGN KEY (author_id) REFERENCES user(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// groups
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS groups (
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		author_id INTEGER NOT NULL,
		cover TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (author_id) REFERENCES user(id) ON DELETE CASCADE
	)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// groups members
	_, err = instanceOfDb.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	}

	// groups post
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS groupPosts (
		id INTEGER PRIMARY KEY,
		group_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		image TEXT DEFAULT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
		FOREIGN KEY (author_id) REFERENCES user(id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// groups comments
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS groupComments (
		id INTEGER PRIMARY KEY,
		group_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		image TEXT DEFAULT NULL,
		post_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
		FOREIGN KEY (author_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES groupPosts(id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the "Message Groups" table exists
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS groupMessage (
		id INTEGER PRIMARY KEY,
		sender_id INT,
		group_id INT, 
		content TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (sender_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	//
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS groupMessageStatus (
		user_id INT,
		group_message_id INT,
		status TEXT CHECK (status IN ('read', 'unread')) DEFAULT 'unread',
		FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
		FOREIGN KEY (group_message_id) REFERENCES groupMessage(id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// groups events
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS groupEvents (
		id INTEGER PRIMARY KEY,
		group_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL, 
		datetime DATE NOT NULL,
		author_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
		FOREIGN KEY (author_id) REFERENCES user(id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// groups events responses
	_, err = instanceOfDb.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	}

	// notification
	_, err = instanceOfDb.Exec(`
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
	)
	`)

	if err != nil {
		log.Fatal(err)
	}

	// Check if the "follow" table exists
	_, err = instanceOfDb.Exec(`
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
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "comment_likes" table
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS comment_like (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		entries_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (entries_id) REFERENCES comments(id),
		FOREIGN KEY (author_id) REFERENCES user(id)
	);
	`)

	if err != nil {
		log.Fatal(err)
	}

	// Create the "post_likes" table
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS post_like (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		entries_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (entries_id) REFERENCES post(id),
		FOREIGN KEY (author_id) REFERENCES user(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "post_dislikes" table
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS post_dislike (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		entries_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (entries_id) REFERENCES post(id),
		FOREIGN KEY (author_id) REFERENCES user(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the "Message" table exists
	_, err = instanceOfDb.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the "comment_dislikes" table
	_, err = instanceOfDb.Exec(`
	CREATE TABLE IF NOT EXISTS comment_dislike (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		entries_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (entries_id) REFERENCES comments(id),
		FOREIGN KEY (author_id) REFERENCES user(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// ...

/*
DropAnTable attempts to drop a specified table from the database schema.
It takes an instance of *sql.DB as a parameter and the name of the table to be dropped.
If the specified table exists, it is dropped; otherwise, no action is taken.
*/
func DropAnTable(db *sql.DB, tableName string) {

	_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName))
	if err != nil {
		log.Fatal(err)
	}
}
