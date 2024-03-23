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
	Default().WithOption(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Debugf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Debugf(format, args...)
}

func Info(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Infof(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Infof(format, args...)
}

func Warn(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Warnf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Warnf(format, args...)
}

func Error(msg string, fields ...zap.Field) {
	Default().WithOption(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func Errorf(format string, args ...interface{}) {
	Default().WithOption(zap.AddCallerSkip(1)).Errorf(format, args...)
}
