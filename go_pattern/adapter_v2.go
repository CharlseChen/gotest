package main

import "fmt"

type Logger interface {
	LogInfo(msg string)
	LogError(msg string)
}

type OldLogger struct {
}

func (o *OldLogger) WriteLog(level int, message string) {
	// level: 1=info, 2=error
	fmt.Printf("[Level %d] %s\n", level, message)
}

type OldLoggerAdapter struct {
	OldLogger *OldLogger
}

func (a *OldLoggerAdapter) LogInfo(msg string) {
	a.OldLogger.WriteLog(1, msg)
}

func (a *OldLoggerAdapter) LogError(msg string) {
	a.OldLogger.WriteLog(2, msg)
}

func main() {
	var logger Logger
	logger = &OldLoggerAdapter{new(OldLogger)}
	logger.LogInfo("hello")
}
