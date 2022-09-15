package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

func New(level string, stage string) *Logger {
	var globalLevel zerolog.Level

	switch strings.ToLower(level) {
	case "debug":
		globalLevel = zerolog.DebugLevel
	case "info":
		globalLevel = zerolog.InfoLevel
	case "warn":
		globalLevel = zerolog.WarnLevel
	case "error":
		globalLevel = zerolog.ErrorLevel
	default:
		globalLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(globalLevel)

	var logger zerolog.Logger
	skipFrameCount := 3

	if stage == "dev" {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}

		logger = zerolog.New(output).
			With().
			Timestamp().
			Caller().
			Logger()
	} else {
		logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).
			Logger()
	}

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.log("debug", message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log("info", message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.log("warn", message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.log("error", message, args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.log("fatal", message, args...)
}

func (l *Logger) log(level, message interface{}, args ...interface{}) {

	msg := l.msgFmt(message)

	switch level {
	case "debug":
		l.logger.Debug().Msgf(msg, args...)
	case "info":
		if len(args) == 0 {
			l.logger.Info().Msg(msg)
		} else {
			l.logger.Info().Msgf(msg, args...)
		}
	case "warn":
		if len(args) == 0 {
			l.logger.Warn().Msg(msg)
		} else {
			l.logger.Warn().Msgf(msg, args...)
		}
	case "error":
		if l.logger.GetLevel() == zerolog.DebugLevel {
			l.logger.Debug().Msgf(msg, args...)
		} else {
			l.logger.Error().Msgf(msg, args...)
		}
	case "fatal":
		l.logger.Fatal().Msgf(msg, args...)
	default:
		l.logger.Info().Msg(msg)
	}

}

func (l *Logger) msgFmt(message interface{}) string {
	switch msg := message.(type) {
	case error:
		return msg.Error()
	case string:
		return msg
	default:
		return fmt.Sprintf("%s message has unknown type", message)
	}
}
