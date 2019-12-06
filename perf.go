package perf

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const _EnvPath = "WODADEHENCOU_PERF"
const _Enable = "enable"
const _True = "true"

var enable bool

func init() {
	cfg := os.Getenv(_EnvPath)
	cfg = strings.ToLower(cfg)
	if strings.Compare(cfg, _Enable) == 0 {
		enable = true
		return
	}
	if strings.Compare(cfg, _True) == 0 {
		enable = true
		return
	}
	enable = false
	return
}

func Enable() bool {
	return enable
}

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

func Run(f func(), name string) {
	start := time.Now()
	f()
	diff := time.Since(start).Nanoseconds()
	n := (int64(time.Second) / diff)
	n = n + n/5
	t := 0
	for n > 10 {
		n /= 10
		t++
	}
	for t > 0 {
		n = n * 10
		t--
	}
	fmt.Printf("%s run %d times\n", name, int(n))
	RunNTimes(f, int(n), name)
}

func MayRun(f func(), name string) {
	if Enable() {
		Run(f, name)
	}
}
