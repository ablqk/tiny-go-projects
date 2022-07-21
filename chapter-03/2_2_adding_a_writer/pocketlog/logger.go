package pocketlog

import (
	"fmt"
	"io"
)

// Logger is used to log information.
type Logger struct {
	level  Level
	output io.Writer
}

// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(level Level, output io.Writer) *Logger {
	return &Logger{
		level:  level,
		output: output,
	}
}

// Debug formats and prints a message if the log level is debug or higher.
func (l Logger) Debug(format string, args ...any) {
	if l.level <= LevelDebug {
		_, _ = fmt.Fprintf(l.output, format+"\n", args...)
	}
}

// Info formats and prints a message if the log level is info or higher.
func (l Logger) Info(format string, args ...any) {
	if l.level <= LevelInfo {
		_, _ = fmt.Fprintf(l.output, format+"\n", args...)
	}
}

// Error formats and prints a message if the log level is error or higher.
func (l Logger) Error(format string, args ...any) {
	if l.level <= LevelError {
		_, _ = fmt.Fprintf(l.output, format+"\n", args...)
	}
}
