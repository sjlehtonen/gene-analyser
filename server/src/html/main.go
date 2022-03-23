package html

import (
	"fmt"
	"strings"
)

type elements struct{}
type Html = string

type TableData struct {
	ColumnName string
	Data       []string
}

var Elements = elements{}

func tag(tag string, text ...string) Html {
	return fmt.Sprintf("<%s>%s</%s>", tag, strings.Join(text, ""), tag)
}

func (elements *elements) P(text ...string) Html {
	return tag("p", text...)
}

func (elements *elements) Style(targets []string, style string) Html {
	return "<style>" + strings.Join(targets, ", ") + "{ " + style + "}" + "</style>"
}

func (elements *elements) Table(data []TableData) Html {
	finalHtml := "<table>"
	headerHtml := "<tr>"
	for _, v := range data {
		headerHtml += "<th>" + v.ColumnName + "</th>"
	}
	headerHtml += "</tr>"
	rowsHtml := ""
	rowCount := len(data[0].Data)

	for i := 0; i < rowCount; i += 1 {
		rowsHtml += "<tr>"
		for _, v := range data {
			rowsHtml += "<td>" + fmt.Sprint(v.Data[i]) + "</td>"
		}
		rowsHtml += "</tr>"
	}
	finalHtml += headerHtml
	finalHtml += rowsHtml
	finalHtml += "</table>"
	return finalHtml
}

func (elements *elements) Html(content ...string) Html {
	return tag("html", content...)
}

func (elements *elements) B(text ...string) Html {
	return tag("b", text...)
}

func (elements *elements) Div(text ...string) Html {
	return tag("div", text...)
}
