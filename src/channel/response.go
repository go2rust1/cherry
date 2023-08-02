package channel

import "github.com/go2rust1/cherry/src/trait"

var _response = make(chan trait.Response)

func ResponseIn(response trait.Response) {
	go func() { _response <- response }()
}

func ResponseChan() chan trait.Response {
	return _response
}
