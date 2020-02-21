package jth

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nsf/termbox-go"
)

type Markup struct {
	Foreground   termbox.Attribute
	Background   termbox.Attribute
	RightAligned bool
	tags         map[string]termbox.Attribute
	regex        *regexp.Regexp
}

func NewMarkup() *Markup {
	markup := &Markup{}
	markup.Foreground = termbox.ColorDefault
	markup.Background = termbox.ColorDefault
	markup.RightAligned = false

	markup.tags = make(map[string]termbox.Attribute)
	markup.tags[`/`] = termbox.ColorDefault
	markup.tags[`black`] = termbox.ColorBlack
	markup.tags[`red`] = termbox.ColorRed
	markup.tags[`green`] = termbox.ColorGreen
	markup.tags[`yellow`] = termbox.ColorYellow
	markup.tags[`blue`] = termbox.ColorBlue
	markup.tags[`magenta`] = termbox.ColorMagenta
	markup.tags[`cyan`] = termbox.ColorCyan
	markup.tags[`white`] = termbox.ColorWhite
	markup.tags[`right`] = termbox.ColorDefault
	markup.tags[`b`] = termbox.AttrBold
	markup.tags[`u`] = termbox.AttrUnderline
	markup.tags[`r`] = termbox.AttrReverse
	markup.regex = markup.supportedTags()

	return markup
}

func (m *Markup) Tokenize(s string) []string {
	fmt.Println(s)
	matches := m.regex.FindAllStringIndex(s, -1)
	strings := make([]string, 0, len(matches))

	head, tail := 0, 0
	for _, match := range matches {
		tail = match[0]
		if match[1] != 0 {
			if head != 0 || tail != 0 {
				strings = append(strings, s[head:tail])
			}

			strings = append(strings, s[match[0]:match[1]])

		}
		head = match[1]
	}

	if head != len(s) && tail != len(s) {
		strings = append(strings, s[head:])
	}

	return strings
}

func (m *Markup) IsTag(s string) bool {
	tag, open := probeForTag(s)

	if tag == `` {
		return false
	}

	return m.process(tag, open)
}

func (m *Markup) process(tag string, open bool) bool {
	if attr, ok := m.tags[tag]; ok {
		switch tag {
		case `right`:
			m.RightAligned = open
		default:
			if open {
				if attr >= termbox.AttrBold {
					m.Foreground |= attr
				} else {
					m.Foreground = attr
				}
			} else {
				if attr >= termbox.AttrBold {
					m.Foreground &= ^attr
				} else {
					m.Foreground = termbox.ColorDefault
				}
			}
		}
	}

	return true
}

func (m *Markup) supportedTags() *regexp.Regexp {
	arr := []string{}

	for tag := range m.tags {
		arr = append(arr, `</?`+tag+`>`)
	}

	return regexp.MustCompile(strings.Join(arr, `|`))
}

func probeForTag(s string) (string, bool) {
	if len(s) > 2 && s[0:1] == `<` && s[len(s)-1:] == `>` {
		return extractTagName(s), s[1:2] != `/`
	}

	return ``, false
}

func extractTagName(s string) string {
	if len(s) < 3 {
		return ``
	} else if s[1:2] != `/` {
		return s[1 : len(s)-1]
	} else if len(s) > 3 {
		return s[2 : len(s)-1]
	}

	return `/`
}
