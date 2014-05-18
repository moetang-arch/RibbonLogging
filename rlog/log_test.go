package rlog

import (
	"fmt"
	"testing"
)

func Test_Logger(t *testing.T) {
	//logger -> format -> appender
	//%date %level [%thread] %logger{10} [%file:%line] %msg%n
	//%contextName [%t] %d
	//properties - ${USER_HOME} , etc.
	logFab := CreateLogFactory()
	logFab.CreateLoggerWithDefaultConfig("theTestLogger")
	logger := logFab.GetLogger("hello")
	fmt.Println(fmt.Sprint("hello"))
	logger.Trace("trace")
	logger.Debug("haha %s", "ljljlklk")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.Fatal("fatal")
	logger.Panic("panic")

	logger2 := logFab.GetLogger("theTestLogger")
	logger2.Trace("trace")
}
