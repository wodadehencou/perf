package perf

import (
	"testing"
	"time"
)

func Test_Run(t *testing.T) {
	Run(func() {
		return
	}, "dummy")
	Run(func() {
		time.Sleep(1 * time.Millisecond)
		return
	}, "ms")
}
