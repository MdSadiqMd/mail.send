package main

import (
	"fmt"
	"time"

	logger "github.com/MdSadiqMd/mail.send.git/pkg/log"
)

func main() {
	apiLogger := logger.New("api-server")
	dbLogger := logger.New("database")
	authLogger := logger.New("auth-service")
	cacheLogger := logger.New("cache")
	webLogger := logger.New("web-frontend")
	queueLogger := logger.New("message-queue")
	demonstrateLogging(apiLogger, dbLogger, authLogger, cacheLogger, webLogger, queueLogger)
}

func demonstrateLogging(apiLogger, dbLogger, authLogger, cacheLogger, webLogger, queueLogger *logger.Logger) {
	// API Server logs
	apiLogger.Info("Server starting on port 8080")
	apiLogger.Debug("Loading configuration files")
	apiLogger.Warn("High memory usage detected", "usage: 85%")
	apiLogger.Error("Request failed", "status: 500", "path: /api/users")

	// Database logs
	dbLogger.Info("Connecting to PostgreSQL")
	dbLogger.Debug("Query executed successfully", "duration: 150ms")
	dbLogger.Warn("Connection pool nearly full", "connections: 95/100")
	dbLogger.Error("Query timeout", "query: SELECT * FROM users")

	// Auth Service logs
	authLogger.Info("Authentication service initialized")
	authLogger.Debug("Token validation completed", "user: john_doe")
	authLogger.Warn("Multiple failed login attempts", "ip: 192.168.1.100")
	authLogger.Error("JWT token expired", "user: admin", "expired: 30m ago")

	// Cache logs
	cacheLogger.Info("Redis cache connected")
	cacheLogger.Debug("Cache hit", "key: user_session_123", "ttl: 2h")
	cacheLogger.Warn("Cache memory usage high", "usage: 89%")
	cacheLogger.Error("Cache connection lost", "reconnecting in 5s")

	// Web Frontend logs
	webLogger.Info("Frontend assets compiled")
	webLogger.Debug("Static files served", "path: /assets/app.js")
	webLogger.Warn("Slow page load detected", "page: /dashboard", "time: 3.2s")
	webLogger.Error("JavaScript error", "file: app.js", "line: 245")

	// Message Queue logs
	queueLogger.Info("Message queue initialized")
	queueLogger.Debug("Message processed", "queue: email", "id: msg_123")
	queueLogger.Warn("Queue backlog growing", "pending: 1500", "queue: notifications")
	queueLogger.Error("Failed to process message", "queue: payments", "error: timeout")

	// Demonstrate all log levels in sequence
	time.Sleep(500 * time.Millisecond)

	apiLogger.Debug("Debug message from API")
	dbLogger.Info("Info message from Database")
	authLogger.Warn("Warning from Auth Service")
	cacheLogger.Error("Error from Cache")

	// Show more services to demonstrate color variety
	for i := 1; i <= 5; i++ {
		serviceName := fmt.Sprintf("service-%d", i)
		tempLogger := logger.New(serviceName)
		tempLogger.Info("Service initialized", "instance:", i)
		time.Sleep(100 * time.Millisecond)
	}
}
