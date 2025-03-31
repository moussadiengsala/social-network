package database

import "database/sql"

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTables() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(createUserTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createPostTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createCommentTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createCategoryTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createPrivateMessageTable)
	if err != nil {
		return err
	}
	return nil
}
