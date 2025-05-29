package logger

import (
	"fmt"
	"hash/fnv"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type Logger struct {
	serviceName  string
	serviceStyle lipgloss.Style
	logger       *log.Logger
}

var (
	serviceColors = []string{
		"10", // Bright Green
		"13", // Bright Magenta
		"14", // Bright Cyan
		"11", // Bright Yellow
		"9",  // Bright Red
		"12", // Bright Blue
		"5",  // Magenta
		"6",  // Cyan
		"3",  // Yellow
		"4",  // Blue
		"2",  // Green
		"1",  // Red
		"8",  // Gray
		"7",  // Light Gray
		"15", // White
	}
)

func New(serviceName string) *Logger {
	logger := log.New(os.Stdout)
	styles := log.DefaultStyles()
	logger.SetStyles(styles)
	logger.SetReportTimestamp(false)
	logger.SetReportCaller(false)

	return &Logger{
		serviceName:  serviceName,
		serviceStyle: getServiceStyle(serviceName),
		logger:       logger,
	}
}

func getServiceStyle(serviceName string) lipgloss.Style {
	hash := fnv.New32a()
	hash.Write([]byte(serviceName))
	colorIndex := int(hash.Sum32()) % len(serviceColors)
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(serviceColors[colorIndex])).
		Bold(true)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.logWithService(log.DebugLevel, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.logWithService(log.InfoLevel, format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.logWithService(log.WarnLevel, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.logWithService(log.ErrorLevel, format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	var message string
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	} else {
		message = format
	}
	l.logger.Fatal(message)
}

func (l *Logger) logWithService(level log.Level, format string, args ...interface{}) {
	styles := log.DefaultStyles()
	levelStyle := styles.Levels[level]
	var levelText string
	switch level {
	case log.DebugLevel:
		levelText = "DEBU"
	case log.InfoLevel:
		levelText = "INFO"
	case log.WarnLevel:
		levelText = "WARN"
	case log.ErrorLevel:
		levelText = "ERRO"
	case log.FatalLevel:
		levelText = "FATA"
	}

	levelRendered := levelStyle.Render(levelText)
	serviceRendered := l.serviceStyle.Render(l.serviceName)

	var message string
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	} else {
		message = format
	}

	fmt.Printf("%s %s %s\n",
		levelRendered,
		serviceRendered,
		message,
	)
}
