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

type CleanText string

func cleanSql(s string) string {
	c := s
	for _, f := range forbidden {
		c = strings.ReplaceAll(c, f, "")
	}
	return c
}

func CleanSql(s string) CleanText {
	return CleanText(cleanSql(s))
}

func Clean(s string) CleanText {
	c := cleanSql(s)
	c = html.EscapeString(c)
	return CleanText(strings.TrimSpace(c))
}
