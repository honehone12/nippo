package input

import (
	"html"
	"nippo/schema"
	"strings"
)

var forbidden = []string{
	"*", "%",
	"`", `\`,
	"--", "#",
	";", "=",
	"OR", "AND", "UNION",
}

func cleanSql(s string) string {
	c := s
	for _, f := range forbidden {
		c = strings.ReplaceAll(c, f, "")
	}
	return c
}

func CleanSql(s string) schema.CleanText {
	return schema.CleanText(cleanSql(s))
}

func Clean(s string) schema.CleanText {
	c := cleanSql(s)
	c = html.EscapeString(c)
	return schema.CleanText(strings.TrimSpace(c))
}
