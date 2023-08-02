package main

import (
	"encoding/json"
	"net/http"
	"github.com/go2rust1/cherry"
)

// example0.go案例说明:
// 一张数据库表绑定一个模型
// 一个主题绑定一个数据库表

type Company struct {
	Name              string `db:"NAME"`
	IsListed          bool   `db:"IS_LISTED"`
	RegisteredCapital int    `db:"REGISTER_CAPITAL"`
	RegisteredAddress string `db:"REGISTER_ADDRESS"`
}

type Frontend struct {
	Name string `json:"name"`
}

func Parser1(topic cherry.Topic, response cherry.Response) {
	// 如果网页返回的数据是json格式, 则使用response.Body()获取字节流
	body := response.Body()

	var front Frontend
	if err := json.Unmarshal(body, &front); err != nil {
		return
	}

	// 构造自定义HTTP请求
	request, _ := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)

	// 发起二级请求, 二级请求使用Parser2解析, 并将该层获取的信息name透传到Parser2
	topic.Requests(request, Parser2, map[string]interface{}{"name": front.Name})
}

func Parser2(topic cherry.Topic, response cherry.Response) {
	// 如果网页返回的数据是html格式, 则使用response.Text()获取字符串
	text := response.Text()
	_ = text
	
	// 假设从text字符串中解析出了以下信息
	isListed := true
	registeredCapital := 10000
	registeredAddress := "北京"

	// 从Parser1解析器传过来的数据
	name := response.Meta()["name"].(string)

	company := Company{
		Name:              name,
		IsListed:          isListed,
		RegisteredCapital: registeredCapital,
		RegisteredAddress: registeredAddress,
	}

	// 将数据发送到后台数据库
	topic.Send(company)
}

func main() {
	// 新建一个MySQL实例
	var db = cherry.MySQL()
	// 设置data source name
	db.SetDSN("root:root@tcp(127.0.0.1:3306)/cherry?parseTime=true&loc=Local")
	// 设置表名
	db.SetTable("company")
	// 绑定Company模型, 模型的标签必须是是表字段的子集
	db.Bind(Company{})

	// 新建任务
	_cherry := cherry.New()

	// 新建一个名为企业的主题
	company, _ := _cherry.NewTopic("企业")
	// 设置起始爬取链接
	company.Request("https://www.baidu.com", Parser1, nil)
	// 主题跟数据库绑定
	company.Bind(db)

	// 启动任务
	_cherry.Start()
}
