package perf

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// RunNTimes is a function to run f() n times, return one time of f()
func RunNTimes(f func(), n int, name string) {
	start := time.Now()
	for i := 0; i < n; i++ {
		f()
	}
	diff := time.Since(start).Nanoseconds()
	if name == "" {
		name = runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	}
	fmt.Printf("==== One time %s is about %s\n", name, DisplayTime(diff/int64(n)))
}

// DisplayTime display int64 ns time in proper format
func DisplayTime(t int64) string {
	var unit string
	var d int64
	if t < 10000 {
		unit = "ns"
		d = t
	} else if t < 10000000 {
		unit = "us"
		d = t / 1000
	} else if t < 10000000000 {
		unit = "ms"
		d = t / 1000000
	} else {
		unit = "s"
		d = t / 1000000000
	}
	return fmt.Sprintf("%d %s", d, unit)
}
