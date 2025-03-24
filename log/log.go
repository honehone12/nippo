package log

import "github.com/sirupsen/logrus"

var L *logrus.Logger

func init() {
	L = logrus.New()
}

func DebugMode(level logrus.Level) {
	L.Formatter = new(logrus.TextFormatter)
	L.Formatter.(*logrus.TextFormatter).FullTimestamp = true
	L.Level = level
}
