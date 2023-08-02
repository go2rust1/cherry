package downloader

import (
	"crypto/tls"
	"io"
	"net/http"
	"github.com/go2rust1/cherry/src/channel"
	"github.com/go2rust1/cherry/src/conf"
	"github.com/go2rust1/cherry/src/counter"
	"github.com/go2rust1/cherry/src/limiter"
	"github.com/go2rust1/cherry/src/logger"
	"github.com/go2rust1/cherry/src/resource"
	"github.com/go2rust1/cherry/src/trait"
	"go.uber.org/zap"
)

func Download(request trait.Request) {
	counter.Increase()
	defer counter.Decrease()
	defer func() { <-limiter.RequestLimiter }()
	client := &http.Client{
		Timeout: conf.RequestTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	responds, err := client.Do(request.Request())
	if err == nil {
		body, _ := io.ReadAll(responds.Body)
		channel.ResponseIn(resource.NewResponse(request, body))
		_ = responds.Body.Close()
		return
	}
	if overflow(request.Finger()) {
		logger.Logger.Error(
			"fail to fetch url",
			zap.String("topic", request.Topic().Name()),
			zap.String("url", request.Request().URL.String()),
			zap.Int("attempt", conf.RequestRetries),
		)
		return
	}
	channel.RequestIn(request)
}
