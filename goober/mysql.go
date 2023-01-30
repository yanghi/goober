package goober

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlRowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
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

type rowTransformer = func(key string, val any) any

func mysqlRowsTranform(rows *sql.Rows, transformer rowTransformer) ([]map[string]interface{}, error) {
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
			v = transformer(col, val)

			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	return tableData, nil
}

func FormatMap(m *map[string]interface{}, format map[string]string) *map[string]interface{} {

	return m
}

type Formatter struct {
	Type string
	Key  string
}

func toFormatter(s string) *Formatter {
	var ss = strings.Split(s, " ")
	var f = &Formatter{}

	for _, v := range ss {
		kvs := strings.Split(v, ":")
		if kvs[0] == "type" {
			f.Type = kvs[1]
		} else if kvs[0] == "key" {
			f.Key = kvs[1]
		}
	}

	return f
}
