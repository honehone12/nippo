package input

import (
	"nippo/schema"
	"strings"
)

var forbidden = strings.NewReplacer(
	"*", "",
	"%", "",
	`"`, "",
	`'`, "",
	"`", "",
	`\`, "",
	"--", "",
	"#", "",
	";", "",
	"=", "",
	"(", "",
	")", "",
	"|", "",
	"&", "",
	"OR", "",
	"AND", "",
	"UNION", "",
	"<", "",
	">", "",
)

func Clean(s string) schema.CleanText {
	return schema.CleanText(strings.TrimSpace(forbidden.Replace(s)))
}
