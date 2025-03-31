package input

import (
	"html"
	"strings"
)

var forbidden = []string{
	"*", "%",
	"`", `\`,
	"--", "#",
	";", "=",
	"OR", "AND", "UNION",
}

func CleanSql(s string) string {
	n := s
	for _, f := range forbidden {
		n = strings.ReplaceAll(n, f, "")
	}
	return n
}

func Clean(s string) string {
	n := CleanSql(s)
	n = html.EscapeString(n)
	return strings.TrimSpace(n)
}
