package rlog

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	blackColor = "\x1b[0m"
)

type appender interface {
	Println(logger *loggerImpl, level Level, logString string)
}

type baseAppender struct {
	traceColor string
	debugColor string
	infoColor  string
	warnColor  string
	errorColor string
	fatalColor string
	panicColor string
}

type consoleAppender struct {
	*baseAppender
}

func (appender *consoleAppender) Println(logger *loggerImpl, level Level, logString string) {
	switch level {
	case TRACE:
		fmt.Println(appender.traceColor, default_TimeString(), loggerInfo(logger), logString)
	case DEBUG:
		fmt.Println(appender.debugColor, default_TimeString(), loggerInfo(logger), logString)
	case INFO:
		fmt.Println(appender.infoColor, default_TimeString(), loggerInfo(logger), logString)
	case WARN:
		fmt.Println(appender.warnColor, default_TimeString(), loggerInfo(logger), logString)
	case ERROR:
		fmt.Println(appender.errorColor, default_TimeString(), loggerInfo(logger), logString)
	case FATAL:
		fmt.Println(appender.fatalColor, default_TimeString(), loggerInfo(logger), logString)
	case PANIC:
		fmt.Println(appender.panicColor, default_TimeString(), loggerInfo(logger), logString)
	}
}

func loggerInfo(logger *loggerImpl) string {
	_, filePath, line, ok := runtime.Caller(4)
	if logger.stackTrace && ok {
		_, file := filepath.Split(filePath)
		return fmt.Sprintf("[%s:%s:%d]", logger.loggerName, file, line)
	} else {
		return fmt.Sprintf("[%s]", logger.loggerName)
	}
}
