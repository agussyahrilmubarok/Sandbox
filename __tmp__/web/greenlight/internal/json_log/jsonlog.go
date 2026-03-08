package json_log

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// `Level` defines the severity of log messages, represented as an integer value.
// The log levels increase in severity from `LevelInfo` to `LevelFatal`.
type Level int8

const (
	LevelInfo  Level = iota // Informational messages
	LevelError              // Error messages that indicate something went wrong
	LevelFatal              // Critical errors that require an application shutdown
	LevelOff                // Disables logging
)

// The `String` method converts a log level into its corresponding string representation.
// It is useful when formatting log entries for output.
func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// `Logger` is a struct that handles logging operations. It includes:
// - `out`: the output destination (e.g., a file or standard output).
// - `minLevel`: the minimum severity level to log.
// - `mu`: a mutex to ensure thread-safe logging when multiple goroutines write concurrently.
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// `New` creates a new `Logger` instance with the provided output destination and minimum logging level.
// This is the constructor function for initializing the logger.
func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

// `print` is the core function responsible for formatting and writing log messages.
// It logs messages that meet or exceed the minimum logging level and formats them as JSON.
func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	// If the log level is below the minimum level, no log is written.
	if level < l.minLevel {
		return 0, nil
	}

	// The log message structure that will be marshaled into JSON.
	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339), // Current time in UTC, formatted according to RFC3339.
		Message:    message,
		Properties: properties,
	}

	// If the log level is `LevelError` or higher, capture the stack trace.
	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	// Marshal the log message struct into JSON.
	var line []byte
	line, err := json.Marshal(aux)
	if err != nil {
		// If marshaling fails, log an error message instead.
		line = []byte(LevelError.String() + ": unable to marshal log message: " + err.Error())
	}

	// Lock to ensure only one goroutine writes to the log at a time.
	l.mu.Lock()
	defer l.mu.Unlock()

	// Write the log entry, appending a newline character.
	return l.out.Write(append(line, '\n'))
}

// `Write` allows the logger to implement the `io.Writer` interface, treating each log as an error.
// It writes error-level log entries.
func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}

// `PrintInfo` logs informational messages at the `LevelInfo` level.
// Properties (metadata) can be added to provide more context for the log.
func (l *Logger) PrintInfo(message string, properties map[string]string) {
	l.print(LevelInfo, message, properties)
}

// `PrintError` logs error messages at the `LevelError` level.
// The error object is converted to a string, and additional properties can be added.
func (l *Logger) PrintError(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
}

// `PrintFatal` logs fatal error messages at the `LevelFatal` level and then exits the application.
// This is used for critical errors that require immediate shutdown.
func (l *Logger) PrintFatal(err error, properties map[string]string) {
	l.print(LevelFatal, err.Error(), properties)
	os.Exit(1)
}
