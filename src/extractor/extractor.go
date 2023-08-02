package extractor

import (
	"github.com/go2rust1/cherry/src/counter"
	"github.com/go2rust1/cherry/src/limiter"
	"github.com/go2rust1/cherry/src/trait"
)

func Extract(response trait.Response) {
	counter.Increase()
	defer counter.Decrease()
	defer func() { <-limiter.ResponseLimiter }()
	response.Request().Parser()(response.Request().Topic(), response)
}
