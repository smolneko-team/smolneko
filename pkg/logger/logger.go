package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Interface interface {
	Debug(message any, args ...any)
	Info(message string, args ...any)
	Warn(message string, args ...any)
	Error(message any, args ...any)
	Fatal(message any, args ...any)
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

	// is the number of stack frames to skip to find the caller
	skipFrameCount := 3

	if stage == "dev" {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		output.FormatLevel = func(i any) string {
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
			CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount). // 5
			Logger()
	}

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) Debug(message any, args ...any) {
	l.log("debug", message, args...)
}

func (l *Logger) Info(message string, args ...any) {
	l.log("info", message, args...)
}

func (l *Logger) Warn(message string, args ...any) {
	l.log("warn", message, args...)
}

func (l *Logger) Error(message any, args ...any) {
	l.log("error", message, args...)
}

func (l *Logger) Fatal(message any, args ...any) {
	l.log("fatal", message, args...)
}

func (l *Logger) log(level, message any, args ...any) {

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

func (l *Logger) msgFmt(message any) string {
	switch msg := message.(type) {
	case error:
		return msg.Error()
	case string:
		return msg
	default:
		return fmt.Sprintf("%s message has unknown type", message)
	}
}
