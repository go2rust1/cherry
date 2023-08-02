package main

import (
	"time"
	"github.com/go2rust1/cherry"
)

// example1.go案例说明:
// 一张数据库表绑定一个模型
// 多个主题绑定一个数据库表

type News struct {
	Url   string    `db:"URL"`
	Title string    `db:"TITLE"`
	Date  time.Time `db:"DATE"`
}

func Parser1(topic cherry.Topic, response cherry.Response) {
	// do your staff
}

func Parser2(topic cherry.Topic, response cherry.Response) {
	// do your staff
}

func Parser3(topic cherry.Topic, response cherry.Response) {
	// do your staff
}

func main() {
	db1 := cherry.DB2()
	db1.SetDSN("DATABASE=cherry;HOSTNAME=localhost;PORT=50000;UID=rust;PWD=rust")
	db1.SetSchema("go2rust")
	db1.SetTable("news_master")
	db1.Bind(News{})

	db2 := cherry.DB2()
	db2.SetDSN("DATABASE=cherry;HOSTNAME=localhost;PORT=50000;UID=rust;PWD=rust")
	db2.SetSchema("go2rust")
	db2.SetTable("news_slaves")
	db2.Bind(News{})

	db3 := cherry.MySQL()
	db3.SetDSN("root:root@tcp(127.0.0.1:3306)/cherry?parseTime=true&loc=Local")
	db3.SetTable("news")
	db3.Bind(News{})

	db4 := cherry.Oracle()
	db4.SetDSN("jason/jason@localhost:1521/helowin")
	db4.SetSchema("jason")
	db4.SetTable("news_nba")
	db4.Bind(News{})

	db5 := cherry.Oracle()
	db5.SetDSN("jason/jason@localhost:1521/helowin")
	db5.SetSchema("jason")
	db5.SetTable("news_cba")
	db5.Bind(News{})

	_cherry := cherry.New()

	topic1, _ := _cherry.NewTopic("主题1")
	topic1.Request("https://www.baidu.com/", Parser1, nil)
	topic1.Bind(db1, db2, db3)

	topic2, _ := _cherry.NewTopic("主题2")
	topic2.Request("https://www.baidu.com/", Parser2, nil)
	topic2.Bind(db4)

	topic3, _ := _cherry.NewTopic("主题3")
	topic3.Request("https://www.baidu.com/", Parser3, nil)
	topic3.Bind(db5)

	_cherry.Start()
}
