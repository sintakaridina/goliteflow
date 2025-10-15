package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger wraps zerolog with additional functionality
type Logger struct {
	logger zerolog.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	// Configure console output with pretty formatting
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	logger := zerolog.New(output).With().Timestamp().Logger()

	return &Logger{
		logger: logger,
	}
}

// NewFileLogger creates a logger that writes to a file
func NewFileLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := zerolog.New(file).With().Timestamp().Logger()

	return &Logger{
		logger: logger,
	}, nil
}

// SetLevel sets the logging level
func (l *Logger) SetLevel(level zerolog.Level) {
	l.logger = l.logger.Level(level)
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Infof logs a formatted info message
func (l *Logger) Infof(format string, v ...interface{}) {
	l.logger.Info().Msgf(format, v...)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)
}

// Debugf logs a formatted debug message
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logger.Debug().Msgf(format, v...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)
}

// Warnf logs a formatted warning message
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logger.Warn().Msgf(format, v...)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.logger.Error().Msg(msg)
}

// Errorf logs a formatted error message
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logger.Error().Msgf(format, v...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}

// Fatalf logs a formatted fatal message and exits
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatal().Msgf(format, v...)
}

// WithField adds a field to the logger
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		logger: l.logger.With().Interface(key, value).Logger(),
	}
}

// WithFields adds multiple fields to the logger
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	logger := l.logger
	for key, value := range fields {
		logger = logger.With().Interface(key, value).Logger()
	}
	return &Logger{logger: logger}
}

// WithWorkflow adds workflow context to the logger
func (l *Logger) WithWorkflow(workflowID string) *Logger {
	return l.WithField("workflow", workflowID)
}

// WithTask adds task context to the logger
func (l *Logger) WithTask(taskID string) *Logger {
	return l.WithField("task", taskID)
}

// WithExecution adds execution context to the logger
func (l *Logger) WithExecution(executionID string) *Logger {
	return l.WithField("execution", executionID)
}

// GetZerologLogger returns the underlying zerolog logger
func (l *Logger) GetZerologLogger() zerolog.Logger {
	return l.logger
}

// Global logger instance
var globalLogger *Logger

// InitGlobalLogger initializes the global logger
func InitGlobalLogger() {
	globalLogger = NewLogger()
}

// InitGlobalFileLogger initializes the global logger with file output
func InitGlobalFileLogger(filename string) error {
	var err error
	globalLogger, err = NewFileLogger(filename)
	return err
}

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() *Logger {
	if globalLogger == nil {
		InitGlobalLogger()
	}
	return globalLogger
}

// SetGlobalLevel sets the global logging level
func SetGlobalLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
	if globalLogger != nil {
		globalLogger.SetLevel(level)
	}
}

// Convenience functions for global logger
func Info(msg string) {
	GetGlobalLogger().Info(msg)
}

func Infof(format string, v ...interface{}) {
	GetGlobalLogger().Infof(format, v...)
}

func Debug(msg string) {
	GetGlobalLogger().Debug(msg)
}

func Debugf(format string, v ...interface{}) {
	GetGlobalLogger().Debugf(format, v...)
}

func Warn(msg string) {
	GetGlobalLogger().Warn(msg)
}

func Warnf(format string, v ...interface{}) {
	GetGlobalLogger().Warnf(format, v...)
}

func Error(msg string) {
	GetGlobalLogger().Error(msg)
}

func Errorf(format string, v ...interface{}) {
	GetGlobalLogger().Errorf(format, v...)
}

func Fatal(msg string) {
	GetGlobalLogger().Fatal(msg)
}

func Fatalf(format string, v ...interface{}) {
	GetGlobalLogger().Fatalf(format, v...)
}
