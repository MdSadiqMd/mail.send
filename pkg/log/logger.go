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

func (l *Logger) Debug(message string, data ...interface{}) {
	l.logWithService(log.DebugLevel, message, data...)
}

func (l *Logger) Info(message string, data ...interface{}) {
	l.logWithService(log.InfoLevel, message, data...)
}

func (l *Logger) Warn(message string, data ...interface{}) {
	l.logWithService(log.WarnLevel, message, data...)
}

func (l *Logger) Error(message string, data ...interface{}) {
	l.logWithService(log.ErrorLevel, message, data...)
}

func (l *Logger) Fatal(message string, data ...interface{}) {
	l.logWithService(log.FatalLevel, message, data...)
}

func (l *Logger) logWithService(level log.Level, message string, data ...interface{}) {
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
	var dataText string
	if len(data) > 0 {
		dataText = fmt.Sprintf(" %v", data)
	}

	fmt.Printf("%s %s %s%s\n",
		levelRendered,
		serviceRendered,
		message,
		dataText,
	)
}
