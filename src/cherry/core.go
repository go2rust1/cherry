package cherry

import (
	"time"
	"github.com/go2rust1/cherry/src/channel"
	"github.com/go2rust1/cherry/src/conf"
	"github.com/go2rust1/cherry/src/counter"
	"github.com/go2rust1/cherry/src/downloader"
	"github.com/go2rust1/cherry/src/extractor"
	"github.com/go2rust1/cherry/src/limiter"
)

// RoundRobin 轮询
func (c *cherry) RoundRobin() {
	for {
		select {
		case request := <-channel.RequestChan():
			limiter.RequestLimiter <- struct{}{}
			go downloader.Download(request)
		case response := <-channel.ResponseChan():
			limiter.ResponseLimiter <- struct{}{}
			go extractor.Extract(response)
		}
	}
}

// HeartBeatDetection 心跳检测
func (c *cherry) HeartBeatDetection() {
	var n int
	for range time.Tick(time.Second) {
		if counter.Count() == 0 {
			n++
		}
		if n >= conf.HeartBeat {
			c.done <- struct{}{}
			return
		}
	}
}
