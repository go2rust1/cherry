package dbms

import (
	"github.com/go2rust1/cherry/src/logger"
	"github.com/go2rust1/cherry/src/mod/database"
	"github.com/go2rust1/cherry/src/mod/reflects"
	"github.com/go2rust1/cherry/src/trait"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var DBMS = &dbms{
	db:   make(map[string][]trait.Database),
	conn: make(map[string]*sqlx.DB),
	tabs: make(map[string][]string),
	node: make(map[string]*node),
}

type dbms struct {
	// db 一个主题可以对应多个库
	// key: topic name
	// val: []trait.DB
	db map[string][]trait.Database

	// conn 数据库连接
	// key: dns
	// val: *sqlx.DB
	conn map[string]*sqlx.DB

	// tabs 一个主题可以对应多张表
	// key: topic name
	// val: []string{dst}
	tabs map[string][]string

	// node 一张表对应一个节点
	// key: dst
	// val: *node
	node map[string]*node
}

func (d *dbms) Initialize() {
	for topic, db := range d.db {
		d.InitializeConn(topic, db...)
		d.InitializeTabs(topic, db...)
		d.InitializeNode(topic, db...)
	}
	for _, node := range d.node {
		go node.wander()
	}
}

func (d *dbms) InitializeConn(topic string, db ...trait.Database) {
	for _, db := range db {
		inst := db.(database.Database)
		if _, ok := d.conn[inst.DSN()]; ok {
			continue
		}
		conn, err := inst.Open()
		if err != nil {
			logger.Logger.Fatal(
				"failed to initialize database connection",
				zap.String("dsn", inst.DSN()),
			)
		}
		d.conn[inst.DSN()] = conn
	}
}

func (d *dbms) InitializeTabs(topic string, db ...trait.Database) {
	for _, db := range db {
		inst := db.(database.Database)
		if _, ok := d.tabs[topic]; !ok {
			d.tabs[topic] = []string{inst.DST()}
			continue
		}
		d.tabs[topic] = append(d.tabs[topic], inst.DST())
	}
}

func (d *dbms) InitializeNode(topic string, db ...trait.Database) {
	for _, db := range db {
		inst := db.(database.Database)
		if _, ok := d.node[inst.DST()]; !ok {
			inst.Share(d.conn[inst.DSN()])
			d.node[inst.DST()] = &node{db: inst, pipeline: make(chan interface{})}
			continue
		}
		// 检查数据库绑定的模型是否一致
		m1 := inst.Tag()
		m2 := d.node[inst.DST()].db.Tag()
		if !reflects.StringSliceEqualFold(m1, m2) {
			logger.Logger.Fatal(
				"inconsistent models bound to the same database table",
				zap.String("dsn:::schema:::table", inst.DST()),
				zap.Any("model1", m1),
				zap.Any("model2", m2),
			)
		}
	}
}

// Close 任务关闭数据库连接
func (d *dbms) Close() {
	for _, db := range d.conn {
		_ = db.Close()
	}
}

func Bind(topic string, db ...trait.Database) {
	DBMS.db[topic] = db
}

func Send(topic string, model interface{}) {
	for _, dst := range DBMS.tabs[topic] {
		DBMS.node[dst].pipeline <- model
	}
}
