package log

import (
	"fmt"
	"runtime"
	"testing"
)

func TestConsoleWriter(t *testing.T) {
	w := &ConsoleWriter{
		ANSIColor: runtime.GOOS != "windows",
	}

	for _, level := range []string{"debug", "info", "warning", "error", "fatal", "panic", "hahaha"} {
		_, err := fmt.Fprintf(w, `{"time":"2019-07-10T05:35:54.277Z","level":"%s","caller":"test.go:42","error":"i am test error","foo":"bar","n":42,"message":"hello json console writer"}`+"\n", level)
		if err != nil {
			t.Errorf("test json console writer error: %+v", err)
		}
	}
}
