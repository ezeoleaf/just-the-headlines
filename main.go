package main

import (
	"database/sql"
	"fmt"

	"github.com/ezeoleaf/just-the-headlines/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nsf/termbox-go"
)

func initDB(locPath string) *sql.DB {
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
	CREATE TABLE IF NOT EXISTS newspaper(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		country VARCHAR
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

func drawNewspaper(db *sql.DB) {
	nC := models.GetNewspapers(db)
	fmt.Println(nC.Newspapers)
	for _, n := range nC.Newspapers {
		fmt.Println(n.ID)
	}

}

func main() {
	db := initDB("storage.db")

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	ctrlxpressed := false
	drawNewspaper(db)
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlS && ctrlxpressed {
				termbox.Sync()
			}
			if ev.Key == termbox.KeyCtrlQ && ctrlxpressed {
				break loop
			}
			if ev.Key == termbox.KeyCtrlX {
				ctrlxpressed = true
			} else {
				ctrlxpressed = false
			}
		}
	}
}
