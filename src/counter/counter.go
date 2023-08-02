package counter

import "github.com/go2rust1/cherry/src/mod/counter"

var _counter = counter.New()

func Count() uint64 {
	return _counter.Count()
}

func Increase() {
	_counter.Increase()
}

func Decrease() {
	_counter.Decrease()
}
