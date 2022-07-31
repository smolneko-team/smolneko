package logger

import (
    "fmt"
    "os"
    "strings"

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

func New(level string) *Logger {
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

    skipFrameCount := 3 // TODO ?
    logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

    return &Logger{
        logger: &logger,
    }
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
    l.msg("debug", message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
    l.log(message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
    l.log(message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
    if l.logger.GetLevel() == zerolog.DebugLevel {
        l.Debug(message, args...)
    }

    l.msg("error", message, args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
    l.msg("fatal", message, args...)

    os.Exit(1) // TODO ?
}

func (l *Logger) log(message string, args ...interface{}) {
    if len(args) == 0 {
        l.logger.Info().Msg(message)
    } else {
        l.logger.Info().Msgf(message, args...)
    }
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
    switch msg := message.(type) {
    case error:
        l.log(msg.Error(), args...)
    case string:
        l.log(msg, args...)
    default:
        l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
    }
}
