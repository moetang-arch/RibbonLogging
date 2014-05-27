package rlog

import (
	"fmt"
)

func formatf(logger *loggerImpl, level Level, format string, args interface{}) {
	s := fmt.Sprintf(format, args.([]interface{})...)
	for _, v := range logger.appender {
		v.Println(logger, level, s)
	}
}

func formatAll(logger *loggerImpl, level Level, args interface{}) {
	s := fmt.Sprint(args.([]interface{})...)
	for _, v := range logger.appender {
		v.Println(logger, level, s)
	}
}
