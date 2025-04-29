package log

import "github.com/sirupsen/logrus"

var l *logrus.Logger

func FmtDefault() {
	l = logrus.New()
	l.Formatter = new(logrus.JSONFormatter)
}

func FmtPlain() {
	l = logrus.New()
	l.Formatter = new(logrus.TextFormatter)
	l.Formatter.(*logrus.TextFormatter).FullTimestamp = true
}

func SetLevel(level logrus.Level) {
	l.Level = level
}

func Debugf(format string, args ...interface{}) {
	if l == nil {
		return
	}

	l.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	if l == nil {
		return
	}

	l.Debug(args...)
}

func Infof(format string, args ...interface{}) {
	if l == nil {
		return
	}

	l.Infof(format, args...)
}

func Info(args ...interface{}) {
	if l == nil {
		return
	}

	l.Info(args...)
}

func Warnf(format string, args ...interface{}) {
	if l == nil {
		return
	}

	l.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	if l == nil {
		return
	}

	l.Warn(args...)
}

func Errorf(format string, args ...interface{}) {
	if l == nil {
		return
	}

	l.Errorf(format, args...)
}

func Error(args ...interface{}) {
	if l == nil {
		return
	}

	l.Error(args...)
}

func Fatalf(format string, args ...interface{}) {
	if l == nil {
		return
	}

	l.Fatalf(format, args...)
}

func Fatal(args ...interface{}) {
	if l == nil {
		return
	}

	l.Fatal(args...)
}
