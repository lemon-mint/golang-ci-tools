package table

import (
	"bytes"
	"strings"
)

func PipeEscape(s string) string {
	// Escape the pipe character.
	// '|' ==> '\|'
	s = strings.ReplaceAll(s, "|", "\\|")
	return s
}

func NewTable(header []string, rows [][]string) []byte {
	// Generate GFM table.

	//Header
	var buf bytes.Buffer
	buf.WriteString("| ")
	for i, field := range header {
		buf.WriteString(PipeEscape(field))
		if i < len(header)-1 {
			buf.WriteString(" | ")
		}
	}
	buf.WriteString(" |\n")

	//seperator
	buf.WriteString("| ")
	for i := 0; i < len(header); i++ {
		buf.WriteString("---")
		if i < len(header)-1 {
			buf.WriteString(" | ")
		}
	}
	buf.WriteString(" |\n")

	//Rows
	for _, row := range rows {
		buf.WriteString("| ")
		for i, field := range row {
			buf.WriteString(PipeEscape(field))
			if i < len(row)-1 {
				buf.WriteString(" | ")
			}
		}
		buf.WriteString(" |\n")
	}
	return buf.Bytes()
}
