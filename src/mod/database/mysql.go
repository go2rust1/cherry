package database

import (
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go2rust1/cherry/src/counter"
	"github.com/go2rust1/cherry/src/logger"
	"github.com/go2rust1/cherry/src/mod/reflects"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type mysql struct {
	conn   *sqlx.DB
	dsn    string
	schema string
	table  string
	tag    []string
	query  string
}

func (db *mysql) SetDSN(dsn string) {
	db.dsn = dsn
}

func (db *mysql) SetSchema(schema string) {
	db.schema = schema
}

func (db *mysql) SetTable(table string) {
	db.table = table
}

func (db *mysql) Bind(model interface{}) {
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

// StaticQuery 构造MySQL静态查询字符串
// INSERT INTO table (F1,F2,...FN) VALUES (:F1,:F2,...:FN)
func (db *mysql) StaticQuery() {
	field := strings.Join(db.tag, ",")
	place := strings.Join(db.tag, ",:")
	db.query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (:%s)", db.table, field, place)
}

func (db *mysql) DSN() string {
	return db.dsn
}

func (db *mysql) DST() string {
	return db.dsn + ":::" + db.schema + ":::" + db.table
}

func (db *mysql) Tag() []string {
	return db.tag
}

func (db *mysql) Open() (*sqlx.DB, error) {
	return open("mysql", db.dsn)
}

func (db *mysql) Share(conn *sqlx.DB) {
	db.conn = conn
}

func (db *mysql) Insert(model ...interface{}) {
	counter.Increase()
	defer counter.Decrease()
	if _, err := db.conn.NamedExec(db.query, model); err != nil {
		logger.Logger.Error(
			"insert database error",
			zap.String("kind", "mysql"),
			zap.String("dsn", db.dsn),
			zap.String("schema", db.schema),
			zap.String("table", db.table),
			zap.String("error", err.Error()),
		)
	}
}

func NewMySQL() Database {
	return &mysql{}
}
