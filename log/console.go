package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type consoleLogger struct {
	logger zerolog.Logger
}

// NewConsoleLogger factory
func NewConsoleLogger(lvl uint8) Logger {
	return newDefaultConsoleLogger(lvl)
}

func newDefaultConsoleLogger(lvl uint8) Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s:", i))
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	logger := zerolog.New(output).With().Timestamp().Logger()
	if lvl >= 0 && lvl <= 7 {
		logger = logger.Level(zerolog.Level(lvl))
	} else {
		logger = logger.Level(zerolog.InfoLevel)
	}

	return &consoleLogger{logger: logger}
}

func (l consoleLogger) Print(v ...interface{}) {
	l.logger.Print(v...)
}

func (l consoleLogger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l consoleLogger) Println(v ...interface{}) {
	l.logger.Print(fmt.Sprintln(v...))
}

func (l consoleLogger) Fatal(v ...interface{}) {
	l.logger.Fatal().Msg(fmt.Sprint(v...))
}

func (l consoleLogger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatal().Msg(fmt.Sprintf(format, v...))
}

func (l consoleLogger) Fatalln(v ...interface{}) {
	l.logger.Fatal().Msg(fmt.Sprintln(v...))
}

func (l consoleLogger) Panic(v ...interface{}) {
	l.logger.Panic().Msg(fmt.Sprint(v...))
}

func (l consoleLogger) Panicf(format string, v ...interface{}) {
	l.logger.Panic().Msg(fmt.Sprintf(format, v...))
}

func (l consoleLogger) Panicln(v ...interface{}) {
	l.logger.Panic().Msg(fmt.Sprintln(v...))
}

func (l consoleLogger) Debug(v ...interface{}) {
	l.logger.Debug().Msg(fmt.Sprint(v...))
}

func (l consoleLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debug().Msg(fmt.Sprintf(format, v...))
}

func (l consoleLogger) Info(v ...interface{}) {
	l.logger.Info().Msg(fmt.Sprint(v...))
}

func (l consoleLogger) Infof(format string, v ...interface{}) {
	l.logger.Info().Msg(fmt.Sprintf(format, v...))
}

func (l consoleLogger) Warn(v ...interface{}) {
	l.logger.Warn().Msg(fmt.Sprint(v...))
}

func (l consoleLogger) Warnf(format string, v ...interface{}) {
	l.logger.Warn().Msg(fmt.Sprintf(format, v...))
}

func (l consoleLogger) Error(v ...interface{}) {
	l.logger.Error().Msg(fmt.Sprint(v...))
}

func (l consoleLogger) Errorf(format string, v ...interface{}) {
	l.logger.Error().Msg(fmt.Sprintf(format, v...))
}

func (l consoleLogger) WithFields(fields map[string]interface{}) Logger {
	return &consoleLogger{logger: l.logger.With().Fields(fields).Logger()}
}
