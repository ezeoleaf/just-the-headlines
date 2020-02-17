package jth

import (
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

type Screen struct {
	width    int
	height   int
	cleared  bool
	layout   *Layout
	markup   *Markup
	pausedAt *time.Time
}

func NewScreen() *Screen {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	screen := &Screen{}
	screen.layout = NewLayout()
	screen.markup = NewMarkup()

	return screen.Resize()
}

func (screen *Screen) Resize() *Screen {
	screen.width, screen.height = termbox.Size()
	screen.cleared = false

	return screen
}

func (screen *Screen) Pause(pause bool) *Screen {
	if pause {
		screen.pausedAt = new(time.Time)
		*screen.pausedAt = time.Now()
	} else {
		screen.pausedAt = nil
	}

	return screen
}

func (screen *Screen) Clear() *Screen {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	screen.cleared = true

	return screen
}

func (screen *Screen) ClearLine(x int, y int) *Screen {
	for i := x; i < screen.width; i++ {
		termbox.SetCell(i, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Flush()

	return screen
}

func (screen *Screen) Draw(objects ...interface{}) *Screen {
	zonename, _ := time.Now().In(time.Local).Zone()
	if screen.pausedAt != nil {
		defer screen.DrawLine(0, 0, `<right><r>`+screen.pausedAt.Format(`3:04:05pm `+zonename)+`</r></right>`)
	}

	for _, ptr := range objects {
		switch ptr.(type) {
		case *Newspaper:
			object := ptr.(*Newspaper)
			screen.draw(screen.layout.Newspaper(object.Fetch()))
		case *News:
			object := ptr.(*News)
			screen.draw(screen.layout.News(object.Fetch()))
		case time.Time:
			timestamp := ptr.(time.Time).Format(`3:04:05pm ` + zonename)
			screen.DrawLine(0, 0, `<right>`+timestamp+`</right>`)
		default:
			screen.draw(ptr.(string))
		}
	}

	return screen
}

func (screen *Screen) DrawLine(x int, y int, str string) {
	start, column := 0, 0

	for _, token := range screen.markup.Tokenize(str) {
		if screen.markup.IsTag(token) {
			continue
		}

		for i, char := range token {
			if !screen.markup.RightAligned {
				start = x + column
				column++
			} else {
				start = screen.width - len(token) + 1
			}
			termbox.SetCell(x, y, char, screen.markup.Foreground, screen.markup.Background)
		}
	}
	termbox.Flush()
}

func (screen *Screen) draw(str string) {
	if !screen.cleared {
		screen.Clear()
	}
	for row, line := range strings.Split(str, "\n") {
		screen.DrawLine(0, row, line)
	}
}
