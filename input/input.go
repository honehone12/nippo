package input

import (
	"errors"
	"nippo/schema"
	"regexp"
	"strings"
)

const eMAIL_REGEX = `^[A-Za-z0-9._+-]+@[A-Za-z0-9]+\.[A-Za-z]{2,}$`

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
	"<", "",
	">", "",
)
var emailRegex *regexp.Regexp

func Clean(s string) schema.CleanText {
	return schema.CleanText(strings.TrimSpace(forbidden.Replace(s)))
}

func ValidEmail(email string) (schema.CleanText, error) {
	if emailRegex == nil {
		var err error
		if emailRegex, err = regexp.Compile(eMAIL_REGEX); err != nil {
			return "", err
		}
	}

	if strings.Contains(email, "--") {
		return "", errors.New("could not accept email")
	}

	if !emailRegex.MatchString(string(email)) {
		return "", errors.New("invalid email")
	}

	return schema.CleanText(email), nil
}
