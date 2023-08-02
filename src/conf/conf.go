package conf

import "time"

var (
	// Info级别日志文件的存放路径
	InfoFile = "logs/info.log"

	// Info级别日志文件达到多大时触发分割, 单位是MB
	InfoMaxSize = 10

	// 已经被分割的Info级别日志文件最大留存时间, 单位是天
	InfoMaxAge = 1

	// 已经被分割的Info级别日志文件最多留存个数, 单位是个
	InfoMaxBackups = 3

	// 指定被分割的Info级别的日志文件是否要压缩
	InfoCompress = false
)

var (
	// Error及其以上级别日志文件的存放路径
	ErrorFile = "logs/error.log"

	// Error及其以上级别日志文件达到多大时触发分割, 单位是MB
	ErrorMaxSize = 10

	// 已经被分割的Error及其以上级别日志文件最大留存时间, 单位是天
	ErrorMaxAge = 1

	// 已经被分割的Error及其以上级别日志文件最多留存个数, 单位是个
	ErrorMaxBackups = 3

	// 指定被分割的Error及其以上级别日志文件是否要压缩
	ErrorCompress = false
)

var (
	// RequestLimiterNumer 请求限流器并发量
	RequestLimiterNumer = 100

	// ResponseLimiterNumber 响应限流器并发量
	ResponseLimiterNumber = 100
)

var (
	// RequestTimeout 请求下载超时时间
	RequestTimeout = time.Second * 30

	// RequestTimeout 请求下载重试次数
	RequestRetries = 5
)

var (
	// 等待数据超时时间, 如果超过一定时间没有新数据需要插入数据库, 则认为爬虫结束了
	WaitDataTimeout = time.Second * 10
)

var (
	// BatchSize 批量插入数据条数
	BatchSize = 1000
)

var (
	// HeartBeat 心跳次数, 当停止跳动5次, 则程序结束
	HeartBeat = 5
)
