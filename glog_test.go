package log

import (
	"errors"
	"runtime"
	"testing"
	"time"
)

func TestGlogLogger(t *testing.T) {
	log := GlogLogger{
		Level: ParseLevel("debug"),
		Writer: &Writer{
			LocalTime: true,
		},
	}
	log.Infof("bool=%t 1_hour=%s hour_minute_second=%+v, error=%+v, float=%f, int=%d, time=%s",
		true,
		time.Hour,
		[]time.Duration{time.Hour, time.Minute, time.Second},
		errors.New("test error"),
		1.111,
		123456790,
		timeNow())
}

func TestGlogColorLogger(t *testing.T) {
	log := GlogLogger{
		Level:     ParseLevel("debug"),
		ANSIColor: runtime.GOOS != "windows",
		Writer:    &Writer{},
	}
	log.InfoDepth(0, "i am info")
	log.Warningf("i am warning")
	log.Errorf("i am error")
}