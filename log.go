package log

import (
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type Logger interface {
	Log(keyvals ...interface{}) error
}

func Debug(logger Logger, msg string) error {
	return level.Debug(logger).Log("msg", msg)
}

func Debugf(logger Logger, format string, args ...interface{}) error {
	return Debug(logger, fmt.Sprintf(format, args...))
}

func Info(logger Logger, msg string) error {
	return level.Info(logger).Log("msg", msg)
}

func Infof(logger Logger, format string, args ...interface{}) error {
	return Info(logger, fmt.Sprintf(format, args...))
}

func Warn(logger Logger, msg string) error {
	return level.Warn(logger).Log("warn", msg)
}

func Warnf(logger Logger, format string, args ...interface{}) error {
	return Warn(logger, fmt.Sprintf(format, args...))
}

func Error(logger Logger, msg string) error {
	return level.Debug(logger).Log("error", msg)
}

func Errorf(logger Logger, format string, args ...interface{}) error {
	return Error(logger, fmt.Sprintf(format, args...))
}

func With(logger Logger, keyvals ...interface{}) Logger {
	return log.With(logger, keyvals...)
}

type Wrapper struct {
	logger Logger
}

func NewWrapper(logger Logger) *Wrapper {
	return &Wrapper{logger}
}

func (w *Wrapper) Logger() Logger { return w.logger }

func (w *Wrapper) Debug(msg string) error {
	return Debug(w.logger, msg)
}

func (w *Wrapper) Debugf(format string, args ...interface{}) error {
	return Debugf(w.logger, format, args...)
}

func (w *Wrapper) Info(msg string) error {
	return Info(w.logger, msg)
}

func (w *Wrapper) Infof(format string, args ...interface{}) error {
	return Infof(w.logger, format, args...)
}
func (w *Wrapper) Warn(msg string) error {
	return Warn(w.logger, msg)
}

func (w *Wrapper) Warnf(format string, args ...interface{}) error {
	return Warnf(w.logger, format, args...)
}
func (w *Wrapper) Error(msg string) error {
	return Error(w.logger, msg)
}

func (w *Wrapper) Errorf(format string, args ...interface{}) error {
	return Errorf(w.logger, format, args...)
}

func (w *Wrapper) With(keyvals ...interface{}) *Wrapper {
	return NewWrapper(With(w.logger, keyvals...))
}
