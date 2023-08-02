package trait

// Database 暴露给用户的数据库配置项
type Database interface {
	// SetDSN 设置dsn
	SetDSN(string)

	// SetSchema 设置schema
	SetSchema(string)

	// SetTable 设置table
	SetTable(string)

	// Bind 绑定模型
	Bind(interface{})
}
