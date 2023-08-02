package downloader

import (
	"sync"
	"github.com/go2rust1/cherry/src/conf"
)

type tolerance struct {
	sync.Mutex
	finger map[string]int
}

var _tolerance = &tolerance{finger: make(map[string]int)}

func overflow(finger string) bool {
	_tolerance.Mutex.Lock()
	defer _tolerance.Mutex.Unlock()
	n, ok := _tolerance.finger[finger]
	if ok && n >= conf.RequestRetries-1 {
		return true
	}
	_tolerance.finger[finger] += 1
	return false
}
