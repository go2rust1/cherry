package trait

import "net/http"

type Request interface {
	// Topic 获取请求对应的主题 
	Topic() Topic

	// Request 获取原生的HTTP请求
	Request() *http.Request

	// 通过请求携带给响应的页面解析器
	Parser() Parser

	// 通过请求携带给响应的元数据信息
	Meta() Meta

	// 获取请求的指纹信息
	Finger() string
}
