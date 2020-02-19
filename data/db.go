package data

import "database/sql"

func InitDB(locPath string) *sql.DB {
	db, err := sql.Open("sqlite3", locPath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("No DB")
	}

	migrate(db)

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS country(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		code VARCHAR NOT NULL
	);
	CREATE TABLE IF NOT EXISTS newspaper(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		country_id INTEGER REFERENCES country(id)
	);
	CREATE TABLE IF NOT EXISTS section(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		rss VARCHAR NOT NULL,
		failed BOOLEAN DEFAULT FALSE,
		newspaper_id INTEGER REFERENCES newspaper(id)	
	);
	`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
