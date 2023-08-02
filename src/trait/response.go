package trait

// Response 暴露给用户的响应方法
type Response interface {
	// Request 获取响应对应的请求信息 
	Request() Request

	// Body 获取响应数据字节流 
	Body() []byte

	// Text 获取响应数据字符串
	Text() string

	// Meta 获取响应数据元数据
	Meta() Meta
}
