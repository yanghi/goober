package model

import "strings"

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
func formatField(target *map[string]any, field string, formatter Formatter) {

	// var val = target["ss"]

	// if formatter.Key != "" {

	// }

}

func a() {
	var f = make(map[string]any)
	formatField(&f, "ss", Formatter{})
}
