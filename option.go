package logr

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type config struct {
	Level       zap.AtomicLevel
	Format      string
	Options     []zap.Option
	WriteSyncer zapcore.WriteSyncer
}

type Option func(c *config)

func SetLevel(le string) Option {
	return func(c *config) {
		l, err := zapcore.ParseLevel(strings.ToLower(le))
		if err != nil {
			return
		}
		switch l {
		case zapcore.DebugLevel,
			zapcore.InfoLevel,
			zapcore.WarnLevel,
			zapcore.ErrorLevel:
		default:
			return
		}
		c.Level.SetLevel(l)
	}
}

func SetFormat(f string) Option {
	return func(c *config) {
		c.Format = f
	}
}

func SetOptions(opts ...zap.Option) Option {
	return func(c *config) {
		c.Options = opts
	}
}

func SetWriter(ws zapcore.WriteSyncer) Option {
	return func(c *config) {
		c.WriteSyncer = ws
	}
}
