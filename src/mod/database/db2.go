package database

import (
	"fmt"
	"strings"
	_ "github.com/asifjalil/cli"
	"github.com/go2rust1/cherry/src/counter"
	"github.com/go2rust1/cherry/src/logger"
	"github.com/go2rust1/cherry/src/mod/reflects"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type db2 struct {
	conn   *sqlx.DB
	dsn    string
	schema string
	table  string
	tag    []string
	query  string
}

func (db *db2) SetDSN(dsn string) {
	db.dsn = dsn
}

func (db *db2) SetSchema(schema string) {
	db.schema = schema
}

func (db *db2) SetTable(table string) {
	db.table = table
}

// Bind 绑定模型, 将结构体标签绑定到数据库表字段
func (db *db2) Bind(model interface{}) {
	tag, err := reflects.GetStructFieldTags(model, "db")
	if err != nil {
		logger.Logger.Fatal(
			"fail to read model tag",
			zap.Any("model", model),
			zap.Any("error", err.Error()),
		)
	}
	db.tag = tag
	db.StaticQuery()
}

// StaticQuery 构造DB2静态查询字符串
// INSERT INTO schema.table (F1,F2,...FN) VALUES (:F1,:F2,...:FN)
func (db *db2) StaticQuery() {
	field := strings.Join(db.tag, ",")
	place := strings.Join(db.tag, ",:")
	db.query = fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (:%s)", db.schema, db.table, field, place)
}

func (db *db2) DSN() string {
	return db.dsn
}

func (db *db2) DST() string {
	return db.dsn + ":::" + db.schema + ":::" + db.table
}

func (db *db2) Tag() []string {
	return db.tag
}

func (db *db2) Open() (*sqlx.DB, error) {
	return open("cli", db.dsn)
}

func (db *db2) Share(conn *sqlx.DB) {
	db.conn = conn
}

func (db *db2) Insert(model ...interface{}) {
	counter.Increase()
	defer counter.Decrease()
	if _, err := db.conn.NamedExec(db.query, model); err != nil {
		logger.Logger.Error(
			"insert database error",
			zap.String("kind", "db2"),
			zap.String("dsn", db.dsn),
			zap.String("schema", db.schema),
			zap.String("table", db.table),
			zap.String("error", err.Error()),
		)
	}
}

func NewDB2() Database {
	return &db2{}
}
