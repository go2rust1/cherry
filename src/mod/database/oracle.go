package database

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/go2rust1/cherry/src/counter"
	"github.com/go2rust1/cherry/src/logger"
	"github.com/go2rust1/cherry/src/mod/reflects"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type oracle struct {
	conn   *sqlx.DB
	dsn    string
	schema string
	table  string
	tag    []string
}

func (db *oracle) SetDSN(dsn string) {
	db.dsn = dsn
}

func (db *oracle) SetSchema(schema string) {
	db.schema = schema
}

func (db *oracle) SetTable(table string) {
	db.table = table
}

func (db *oracle) Bind(model interface{}) {
	tag, err := reflects.GetStructFieldTags(model, "db")
	if err != nil {
		logger.Logger.Fatal(
			"fail to read model tag",
			zap.Any("model", model),
			zap.Any("error", err.Error()),
		)
	}
	db.tag = tag
}

// _query 构造查询字符串
func (db *oracle) DynamicQuery(model ...interface{}) string {
	query := "INSERT ALL "
	field := strings.Join(db.tag, ",")
	for i := 0; i < len(model); i++ {
		place := strings.Join(db.tag, fmt.Sprintf("%d,:", i))
		query += fmt.Sprintf("INTO %s.%s (%s) VALUES (:%s%d) ", db.schema, db.table, field, place, i)
	}
	query += "SELECT 1 FROM DUAL"
	return query
}

func (db *oracle) DSN() string {
	return db.dsn
}

func (db *oracle) DST() string {
	return db.dsn + ":::" + db.schema + ":::" + db.table
}

func (db *oracle) Tag() []string {
	return db.tag
}

func (db *oracle) Open() (*sqlx.DB, error) {
	return open("godror", db.dsn)
}

func (db *oracle) Share(conn *sqlx.DB) {
	db.conn = conn
}

func (db *oracle) Insert(model ...interface{}) {
	counter.Increase()
	defer counter.Decrease()
	query := db.DynamicQuery(model...)
	if _, err := db.conn.Exec(query, db.DataBuilder(model...)...); err != nil {
		logger.Logger.Error(
			"insert database error",
			zap.String("kind", "oracle"),
			zap.String("dsn", db.dsn),
			zap.String("schema", db.schema),
			zap.String("table", db.table),
			zap.String("error", err.Error()),
		)
	}
}

// DataBuilder 数据构造器
func (db *oracle) DataBuilder(model ...interface{}) (rows []interface{}) {
	for i, row := range model {
		values, _ := reflects.GetStructFieldValues(row, "db")
		for k, v := range values {
			rows = append(rows, sql.Named(fmt.Sprintf("%s%d", k, i), v))
		}
	}
	return
}

func NewOracle() Database {
	return &oracle{}
}
