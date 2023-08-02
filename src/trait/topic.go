package trait

import "net/http"

// Meta 元数据
type Meta = map[string]interface{}

// Parser 响应解析器, 用户根据个人需求从网页源代码中提取对应的数据
type Parser func(Topic, Response)

type Topic interface {
	// Name 主题名称
	Name() string

	// Request 最简单的GET请求, 适用于不带请求头和查询参数的场景
	Request(string, Parser, Meta)

	// Requests 原生的HTTP请求, 适用于用户想构造自定义配置的场景
	Requests(*http.Request, Parser, Meta)

	// Bind 绑定数据库, 一个主题可以绑定多个数据库
	Bind(...Database)

	// Send 将自定义的模型数据发送到绑定的数据库表
	Send(interface{})
}
