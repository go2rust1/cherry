package channel

import "github.com/go2rust1/cherry/src/trait"

var _request = make(chan trait.Request)

func RequestIn(request trait.Request) {
	go func() { _request <- request }()
}

func RequestChan() chan trait.Request {
	return _request
}
