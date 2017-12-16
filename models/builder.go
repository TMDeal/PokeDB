package models

import (
	"bytes"
	"fmt"
	"strings"
)

type Builder interface {
	ToSQL() (string, []interface{})
}

type SelectBuilder struct {
	table       string
	columns     []string
	where       map[string]interface{}
	and         map[string]interface{}
	or          map[string]interface{}
	joins       []string
	limit       int
	limitValid  bool
	offset      int
	offsetValid bool
}

func Select(cols ...string) *SelectBuilder {
	return &SelectBuilder{
		columns: cols,
		where:   make(map[string]interface{}),
		and:     make(map[string]interface{}),
		or:      make(map[string]interface{}),
	}
}

func (s *SelectBuilder) From(table string) *SelectBuilder {
	s.table = table
	return s
}

func (s *SelectBuilder) Where(clause string, value interface{}) *SelectBuilder {
	s.where[clause] = value
	return s
}

func (s *SelectBuilder) Or(clause string, value interface{}) *SelectBuilder {
	s.or[clause] = value
	return s
}

func (s *SelectBuilder) And(clause string, value interface{}) *SelectBuilder {
	s.and[clause] = value
	return s
}

func (s *SelectBuilder) Join(join string) *SelectBuilder {
	s.joins = append(s.joins, "JOIN "+join)
	return s
}

func (s *SelectBuilder) Limit(limit int) *SelectBuilder {
	s.limit = limit
	s.limitValid = true
	return s
}

func (s *SelectBuilder) Offset(offset int) *SelectBuilder {
	s.offset = offset
	s.offsetValid = true
	return s
}

func (s *SelectBuilder) ToSQL() (string, []interface{}) {
	sql := &bytes.Buffer{}
	var args []interface{}

	sql.WriteString("Select ")
	sql.WriteString(fmt.Sprintf("%s FROM ", strings.Join(s.columns, ",")))
	sql.WriteString(s.table + " ")

	if len(s.joins) > 0 {
		for _, join := range s.joins {
			sql.WriteString(fmt.Sprintf("%s ", join))
		}
	}

	if len(s.where) > 0 {
		sql.WriteString("WHERE ")
		for k, v := range s.where {
			sql.WriteString(fmt.Sprintf("%s ", k))
			args = append(args, v)
		}

		if len(s.and) > 0 {
			for k, v := range s.and {
				sql.WriteString(fmt.Sprintf("AND %s ", k))
				args = append(args, v)
			}
		}

		if len(s.or) > 0 {
			for k, v := range s.or {
				sql.WriteString(fmt.Sprintf("OR %s ", k))
				args = append(args, v)
			}
		}
	}

	if s.limitValid {
		sql.WriteString("LIMIT ? ")
		args = append(args, s.limit)
	}

	if s.offsetValid {
		sql.WriteString("OFFSET ? ")
		args = append(args, s.offset)
	}

	return sql.String(), args
}
