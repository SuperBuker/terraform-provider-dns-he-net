package logging

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

func NewZerolog(level zerolog.Level) Logger {
	return &zerologLogger{
		zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02T15:04:05.999Z",
			NoColor:    true,
		}).Level(level).With().Timestamp().Logger(),
	}
}

type zerologLogger struct {
	logger zerolog.Logger
}

func (zerologLogger) withField(e *zerolog.Event, key string, value interface{}) *zerolog.Event {
	switch v := value.(type) {
	case string:
		return e.Str(key, v)
	case int:
		return e.Int(key, v)
	case int8:
		return e.Int8(key, v)
	case int16:
		return e.Int16(key, v)
	case int32:
		return e.Int32(key, v)
	case int64:
		return e.Int64(key, v)
	case uint:
		return e.Uint(key, v)
	case uint8:
		return e.Uint8(key, v)
	case uint16:
		return e.Uint16(key, v)
	case uint32:
		return e.Uint32(key, v)
	case uint64:
		return e.Uint64(key, v)
	case float32:
		return e.Float32(key, v)
	case float64:
		return e.Float64(key, v)
	case bool:
		return e.Bool(key, v)
	case error:
		return e.AnErr(key, v)
	default:
		return e.Interface(key, v)
	}
}

func (l zerologLogger) proc(e *zerolog.Event, msg string, additionalFields ...map[string]interface{}) {
	for _, fields := range additionalFields {
		for key, value := range fields {
			e = l.withField(e, key, value)
		}
	}

	if len(msg) != 0 {
		e.Msg(msg)
	} else {
		e.Send()
	}
}

func (l zerologLogger) Debug(_ context.Context, msg string, additionalFields ...map[string]interface{}) {
	e := l.logger.Debug()

	if e.Enabled() {
		l.proc(e, msg, additionalFields...)
	}
}

func (l zerologLogger) Error(_ context.Context, msg string, additionalFields ...map[string]interface{}) {
	e := l.logger.Error()

	if e.Enabled() {
		l.proc(e, msg, additionalFields...)
	}
}

func (l zerologLogger) Fatal(_ context.Context, msg string, additionalFields ...map[string]interface{}) {
	e := l.logger.Fatal()

	if e.Enabled() {
		l.proc(e, msg, additionalFields...)
	}
}

func (l zerologLogger) Info(_ context.Context, msg string, additionalFields ...map[string]interface{}) {
	e := l.logger.Info()

	if e.Enabled() {
		l.proc(e, msg, additionalFields...)
	}
}

func (l zerologLogger) Panic(_ context.Context, msg string, additionalFields ...map[string]interface{}) {
	e := l.logger.Panic()

	if e.Enabled() {
		l.proc(e, msg, additionalFields...)
	}
}

func (l zerologLogger) Trace(_ context.Context, msg string, additionalFields ...map[string]interface{}) {
	e := l.logger.Trace()

	if e.Enabled() {
		l.proc(e, msg, additionalFields...)
	}
}

func (l zerologLogger) Warn(_ context.Context, msg string, additionalFields ...map[string]interface{}) {
	e := l.logger.Warn()

	if e.Enabled() {
		l.proc(e, msg, additionalFields...)
	}
}
