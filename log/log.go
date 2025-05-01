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

func Tracef(format string, args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Tracef(format, args...)
}

func Trace(args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Trace(args...)
}

func Debugf(format string, args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Debugf(format, args...)
}

func Debug(args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Debug(args...)
}

func Infof(format string, args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Infof(format, args...)
}

func Info(args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Info(args...)
}

func Warnf(format string, args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Warnf(format, args...)
}

func Warn(args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Warn(args...)
}

func Errorf(format string, args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Errorf(format, args...)
}

func Error(args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Error(args...)
}

func Panicf(format string, args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Panicf(format, args...)
}

func Panic(args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Panic(args...)
}

func Fatalf(format string, args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Fatalf(format, args...)
}

func Fatal(args ...any) {
	if l == nil {
		FmtDefault()
	}

	l.Fatal(args...)
}
