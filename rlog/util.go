package rlog

import (
	"time"
)

const (
	timeFormatLayout = "2006-01-02 15:04:05.000"
)

func default_TimeString() string {
	return time.Now().Format(timeFormatLayout)
}
