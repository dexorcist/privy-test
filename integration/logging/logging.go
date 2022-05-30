package logging

import (
	"context"
	"github.com/rs/zerolog"
)

type Logger interface {
	Error(err error, msg string, args ...interface{})
	ErrorWithContext(ctx context.Context, err error, msg string, args ...interface{})
	WarnWithContext(ctx context.Context, msg string, args ...interface{})
	InfoWithContext(ctx context.Context, msg string, args ...interface{})
	DebugWithContext(ctx context.Context, msg string, args ...interface{})
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
}

type logger struct {
	log zerolog.Logger
}

func (l logger) WarnWithContext(ctx context.Context, msg string, args ...interface{}) {
	l.log.Warn().Msgf(msg, args...)
}

func (l logger) InfoWithContext(ctx context.Context, msg string, args ...interface{}) {
	evt := l.log.Info()
	evt.Msgf(msg, args...)
}

func (l logger) DebugWithContext(ctx context.Context, msg string, args ...interface{}) {
	l.log.Debug().Msgf(msg, args...)
}

func (l logger) Warningf(format string, v ...interface{}) {
	l.log.Warn().Msgf(format, v...)
}

func (l logger) Errorf(format string, v ...interface{}) {
	l.log.Error().Msgf(format, v...)
}

func (l logger) Warnf(format string, v ...interface{}) {
	l.log.Warn().Msgf(format, v...)
}

func (l logger) Debugf(format string, v ...interface{}) {
	l.log.Debug().Msgf(format, v...)
}

func (l logger) Error(err error, msg string, args ...interface{}) {
	l.log.Error().Err(err).Msgf(msg, args...)
}

func (l logger) ErrorWithContext(ctx context.Context, err error, msg string, args ...interface{}) {
	l.log.Error().Err(err).Msgf(msg, args...)
}

func NewLog(log zerolog.Logger) Logger {
	return &logger{
		log: log,
	}
}
