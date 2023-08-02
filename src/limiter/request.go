package limiter

import "github.com/go2rust1/cherry/src/conf"

// RequestLimiter 请求限流器
var RequestLimiter = make(chan struct{}, conf.RequestLimiterNumer)
