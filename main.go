package main

import (
	"database/sql"
	"fmt"

	"github.com/ezeoleaf/just-the-headlines/data"
	"github.com/ezeoleaf/just-the-headlines/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nsf/termbox-go"
)

func drawNewspaper(db *sql.DB) {
	nC := models.GetNewspapers(db)
	fmt.Println(nC.Newspapers)
	for _, n := range nC.Newspapers {
		fmt.Println(n.ID)
	}

}

func main() {
	db := data.InitDB("storage.db")

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
