package jth

import (
	"strings"
	"time"

	"github.com/ezeoleaf/just-the-headlines/models"
	"github.com/nsf/termbox-go"
)

type Screen struct {
	width   int
	height  int
	cleared bool
	layout  *Layout
	// markup  *Markup
}

func NewScreen() *Screen {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	s := &Screen{}
	s.layout = NewLayout()

	return s.Resize()
}

func (s *Screen) Close() *Screen {
	termbox.Close()

	return s
}

func (s *Screen) Resize() *Screen {
	s.width, s.height = termbox.Size()
	s.cleared = false

	return s
}

func (s *Screen) Clear() *Screen {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	s.cleared = true

	return s
}

func (s *Screen) ClearLine(x int, y int) *Screen {
	for i := x; i < s.width; i++ {
		termbox.SetCell(i, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Flush()

	return s
}

func (s *Screen) Draw(objects ...interface{}) *Screen {
	zn, _ := time.Now().In(time.Local).Zone()

	for _, ptr := range objects {
		switch ptr.(type) {
		case *models.Newspapers:
			object := ptr.(*models.Newspapers)
			s.draw(s.layout.Newspapers(object.Fetch()))
		case *models.Sections:
			object := ptr.(*models.Sections)
			s.draw(s.layout.Sections(object.Fetch()))
		case *models.News:
			object := ptr.(*models.News)
			s.draw(s.layout.News(object.Fetch()))
		default:
			s.draw(ptr.(string))
		}
	}

	return s
}

func (s *Screen) DrawLine(x int, y int, str string) {
	start, column := 0, 0

	for _, token := range s.markup.Tokenize(str) {
		if s.markup.IsTag(token) {
			continue
		}

		for i, char := range token {
			if !s.markup.RightAligned {
				start = x + column
				column++
			} else {
				start = s.width - len(token) + i
			}

			termbox.SetCell(start, y, char, s.markup.Foreground, s.markup.Background)
		}
	}
	termbox.Flush()
}

func (s *Screen) draw(str string) {
	if !s.cleared {
		s.Clear()
	}

	for r, l := range strings.Split(str, "\n") {
		s.DrawLine(0, r, l)
	}
}
