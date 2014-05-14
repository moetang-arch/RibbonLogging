package rlog

import ()

type Level int8

const (
	TRACE Level = 1
	DEBUG Level = 2
	INFO  Level = 3
	WARN  Level = 4
	ERROR Level = 5
	FATAL Level = 6
	PANIC Level = 7
)

type Logger interface {
	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool
	IsPanicEnabled() bool

	Trace(format string, args ...interface{})
	TraceAll(args ...interface{})
	Debug(format string, args ...interface{})
	DebugAll(args ...interface{})
	Info(format string, args ...interface{})
	InfoAll(args ...interface{})
	Warn(format string, args ...interface{})
	WarnAll(args ...interface{})
	Error(format string, args ...interface{})
	ErrorAll(args ...interface{})
	Fatal(format string, args ...interface{})
	FatalAll(args ...interface{})
	Panic(format string, args ...interface{})
	PanicAll(args ...interface{})
}

func (logger *loggerImpl) IsTraceEnabled() bool {
	return logger.level <= TRACE
}
func (logger *loggerImpl) IsDebugEnabled() bool {
	return logger.level <= DEBUG
}
func (logger *loggerImpl) IsInfoEnabled() bool {
	return logger.level <= INFO
}
func (logger *loggerImpl) IsWarnEnabled() bool {
	return logger.level <= WARN
}
func (logger *loggerImpl) IsErrorEnabled() bool {
	return logger.level <= ERROR
}
func (logger *loggerImpl) IsFatalEnabled() bool {
	return logger.level <= FATAL
}
func (logger *loggerImpl) IsPanicEnabled() bool {
	return logger.level <= PANIC
}

func (logger *loggerImpl) Trace(format string, args ...interface{}) {
	formatf(logger, TRACE, format, args)
}
func (logger *loggerImpl) TraceAll(args ...interface{}) {
	formatAll(logger, TRACE, args)
}
func (logger *loggerImpl) Debug(format string, args ...interface{}) {
	formatf(logger, DEBUG, format, args)
}
func (logger *loggerImpl) DebugAll(args ...interface{}) {
	formatAll(logger, DEBUG, args)
}
func (logger *loggerImpl) Info(format string, args ...interface{}) {
	formatf(logger, INFO, format, args)
}
func (logger *loggerImpl) InfoAll(args ...interface{}) {
	formatAll(logger, INFO, args)
}
func (logger *loggerImpl) Warn(format string, args ...interface{}) {
	formatf(logger, WARN, format, args)
}
func (logger *loggerImpl) WarnAll(args ...interface{}) {
	formatAll(logger, WARN, args)
}
func (logger *loggerImpl) Error(format string, args ...interface{}) {
	formatf(logger, ERROR, format, args)
}
func (logger *loggerImpl) ErrorAll(args ...interface{}) {
	formatAll(logger, ERROR, args)
}
func (logger *loggerImpl) Fatal(format string, args ...interface{}) {
	formatf(logger, FATAL, format, args)
}
func (logger *loggerImpl) FatalAll(args ...interface{}) {
	formatAll(logger, FATAL, args)
}
func (logger *loggerImpl) Panic(format string, args ...interface{}) {
	formatf(logger, PANIC, format, args)
}
func (logger *loggerImpl) PanicAll(args ...interface{}) {
	formatAll(logger, PANIC, args)
}
