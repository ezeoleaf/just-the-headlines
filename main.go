package main

import (
	"github.com/ezeoleaf/just-the-headlines/data"
	"github.com/ezeoleaf/just-the-headlines/jth"
	"github.com/ezeoleaf/just-the-headlines/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nsf/termbox-go"
)

// func drawNewspaper(db *sql.DB) {
// 	nC := models.GetNewspapers(db)
// 	fmt.Println(nC.Newspapers)
// 	for _, n := range nC.Newspapers {
// 		fmt.Println(n.ID)
// 	}

// }

func mainLoop(s *jth.Screen) {
	newspapers := models.NewNewspapers()

	s.Draw(newspapers)
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlS {
				termbox.Sync()
			}
			if ev.Key == termbox.KeyCtrlQ {
				break loop
			}
		}
	}
}

func main() {
	db := data.InitDB("storage.db")
	s := jth.NewScreen(db)
	defer s.Close()

	mainLoop(s)
}
