package jth

import (
	"bytes"
	"text/template"

	"github.com/ezeoleaf/just-the-headlines/models"
)

type Column struct {
	width     int
	name      string
	title     string
	formatter func(...string) string
}

type Layout struct {
	columns []Column
	// sorter            *Sorter
	// filter            *Filter
	newspaperTemplate *template.Template
}

func NewLayout() *Layout {
	layout := &Layout{}
	layout.columns = []Column{
		{10, `ID`, `ID`, nil},
		{10, `Name`, `Name`, nil},
		{10, `Country`, `Country`, nil},
	}

	layout.newspaperTemplate = buildNewspaperTemplate()

	return layout
}

func (l *Layout) Newspapers(newspapers *models.Newspapers) string {
	buffer := new(bytes.Buffer)
	l.newspaperTemplate.Execute(buffer, newspapers)

	return buffer.String()
}

func (l *Layout) TotalColumns() int {
	return len(l.columns)
}

func buildNewspaperTemplate() *template.Template {
	markup := `<right><white>{{.Now}}</></right>



{{.Header}}
{{range.Nespaper}}{{.ID}}{{.Name}}{{.Country}}</>
{{end}}`

	return template.Must(template.New(`newspapers`).Parse(markup))
}