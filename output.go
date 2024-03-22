package logr

import (
	"fmt"
	"go.uber.org/zap"
)

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

func Debug(msg string, fields ...zap.Field) {
	Default().Debug(msg, fields...)
}

func Debugf(format string, args ...interface{}) {
	Default().Debugf(format, args...)
}

func Info(msg string, fields ...zap.Field) {
	Default().Info(msg, fields...)
}

func Infof(format string, args ...interface{}) {
	Default().Infof(format, args...)
}

func Warn(msg string, fields ...zap.Field) {
	Default().Warn(msg, fields...)
}

func Warnf(format string, args ...interface{}) {
	Default().Warnf(format, args...)
}

func Error(msg string, fields ...zap.Field) {
	Default().Error(msg, fields...)
}

func Errorf(format string, args ...interface{}) {
	Default().Errorf(format, args...)
}
