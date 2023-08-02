package trait

type Cherry interface {
	// NewTopic 创建主题, 主题不能重复
	NewTopic(string) (Topic, error)

	// Start 启动爬虫任务
	Start()
}
