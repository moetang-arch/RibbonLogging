package rlog

import (
	// "encoding/xml"
	// "fmt"
	// "io/ioutil"
	"sync"
)

var (
	defaultBaseAppender = &baseAppender{
		traceColor: "\x1b[36m[TRACE]\x1b[0m",
		debugColor: "\x1b[34m[DEBUG]\x1b[0m",
		infoColor:  "\x1b[32m[INFO]\x1b[0m ",
		warnColor:  "\x1b[33m[WARN]\x1b[0m ",
		errorColor: "\x1b[31m[ERROR]\x1b[0m",
		fatalColor: "\x1b[33;41m[FATAL]\x1b[0m",
		panicColor: "\x1b[37;41;5m[PANIC]\x1b[0m",
	}
	defaultLogger = &loggerImpl{
		loggerName: "default",
		level:      TRACE,
		stackTrace: true,
		appender: []appender{
			&consoleAppender{
				baseAppender: defaultBaseAppender,
			},
		},
	}
)

type loggerImpl struct {
	loggerName string
	level      Level
	appender   []appender
	stackTrace bool
}

type LogFactory struct {
	loggerMap   map[string]Logger
	loggerMutex sync.RWMutex
}

func CreateLogFactory() *LogFactory {
	return &LogFactory{
		loggerMap: make(map[string]Logger),
	}
}

func (logFactory *LogFactory) GetLogger(name string) (logger Logger) {
	logFactory.loggerMutex.RLock()
	defer logFactory.loggerMutex.RUnlock()

	logger = logFactory.CreateLoggerWithDefaultConfig(name)
	return
}

func (logFactory *LogFactory) CreateLoggerWithDefaultConfig(name string) (logger Logger) {
	logger, ok := logFactory.loggerMap[name]
	if !ok {
		logger = &loggerImpl{
			loggerName: name,
			level:      TRACE,
			stackTrace: true,
			appender: []appender{
				&consoleAppender{
					baseAppender: defaultBaseAppender,
				},
			},
		}
		logFactory.loggerMap[name] = logger
	}
	return
}

func (logFactory *LogFactory) SetConfigFilePath(path string) {
	// configFilePath = path
}
