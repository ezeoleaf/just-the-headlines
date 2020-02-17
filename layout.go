package jth

import (
	"regexp"
	"text/template"
)

type Column struct {
	width     int
	name      string
	title     string
	formatter func(...string) string
}

type Layout struct {
	columns           []Column
	sorter            *Sorter
	filter            *Filter
	regex             *regexp.Regexp
	newspaperTemplate *template.Template
	newsTemplate      *template.Template
}

func NewLayout() *Layout {
	layout := &Layout{}
	layout.columns = []Column{
		{10, `Newspaper`, `Newspaper`, nil},
	} // Add other columns
	layout.regex = regexp.MustCompile(`(\.\d+)[BMK]?$`)
	layout.newspaperTemplate = buildNewspaperTemplate()
	layout.newsTemplate = buildNewsTemplate()

	return layout
}
