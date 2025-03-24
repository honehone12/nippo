package log

import "github.com/sirupsen/logrus"

var L *logrus.Logger

func FmtDefault() {
	L = logrus.New()
	L.Formatter = new(logrus.JSONFormatter)
}

func FmtPlain() {
	L = logrus.New()
	L.Formatter = new(logrus.TextFormatter)
	L.Formatter.(*logrus.TextFormatter).FullTimestamp = true
}

func SetLevel(level logrus.Level) {
	L.Level = level
}
