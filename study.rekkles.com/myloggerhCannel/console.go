package myloggerChannel

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LogLevel int16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FALTAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)

	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FALTAL, nil
	default:
		err := errors.New("unknow error level")
		return UNKNOWN, err
	}
}

func getLogStrng(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FALTAL:
		return "FALTAL"

	}
	return "DEBUG"

}

type ConsoleLogger struct {
	Level LogLevel
}

func NewConsolelog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{Level: level}
}

func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return c.Level <= LogLevel
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, filename, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogStrng(lv), filename, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {

	c.log(DEBUG, format, a...)

}
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)

}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)

}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)

}
func (c ConsoleLogger) Faltal(format string, a ...interface{}) {
	c.log(FALTAL, format, a...)

}
