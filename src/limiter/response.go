package limiter

import "github.com/go2rust1/cherry/src/conf"

// ResponseLimiter 响应限流器
var ResponseLimiter = make(chan struct{}, conf.ResponseLimiterNumber)
