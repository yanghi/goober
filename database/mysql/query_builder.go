package mysql

import "strings"

type QueryBuilder struct {
	order string
}

func (qb *QueryBuilder) OrderBy(orderBy []string, order string) *QueryBuilder {
	qb.order = qb.BuildOrderBy(orderBy, order)
	return qb
}
func (qb *QueryBuilder) BuildOrderBy(orderBy []string, order string) string {

	if len(orderBy) == 0 {
		return ""
	}

	return "ORDER BY " + strings.Join(orderBy, ",") + " " + order
}
