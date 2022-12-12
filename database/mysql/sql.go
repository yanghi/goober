// mysql
package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConf struct {
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
}

var conf MysqlConf
var DB *sql.DB

// 设置配置
func Config(c MysqlConf) {
	conf = c
}

func New() *sql.DB {
	db, err := sql.Open("mysql", conf.DataSourceName)

	if err != nil {
		fmt.Println("create db err", err)
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	DB = db

	return db
}

// 查询并将结果转换为map
func QueryToMap(db *sql.DB, query string) ([]map[string]interface{}, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	return tableData, nil
}

// 查询并转换为JSON字符串
func QueryToJSON(db *sql.DB, query string) (string, error) {
	m, err := QueryToMap(db, query)
	if err != nil {
		return "", err
	}
	var j, e = json.Marshal(m)

	if e != nil {
		return "", e
	}

	return string(j), nil
}

// todo 配置化
func init() {
	Config(MysqlConf{
		DataSourceName: "root:OLIqMjYR0Gkg6eyJ@/test_blog",
		MaxOpenConns:   10,
		MaxIdleConns:   10,
	})
	fmt.Println("init slq", conf.DataSourceName)
}
