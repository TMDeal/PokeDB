package models

import (
	"bytes"
	"fmt"
)

type Builder interface {
	ToSQL() (string, []interface{})
}

type Conditions struct {
	where       map[string]interface{}
	and         map[string]interface{}
	or          map[string]interface{}
	joins       []string
	limit       int
	limitValid  bool
	offset      int
	offsetValid bool
}

func NewConditions() *Conditions {
	return &Conditions{
		where: make(map[string]interface{}),
		and:   make(map[string]interface{}),
		or:    make(map[string]interface{}),
	}
}

func (s *Conditions) Where(clause string, value interface{}) *Conditions {
	s.where[clause] = value
	return s
}

func (s *Conditions) Or(clause string, value interface{}) *Conditions {
	s.or[clause] = value
	return s
}

func (s *Conditions) And(clause string, value interface{}) *Conditions {
	s.and[clause] = value
	return s
}

func (s *Conditions) Join(join string) *Conditions {
	s.joins = append(s.joins, "JOIN "+join)
	return s
}

func (s *Conditions) Limit(limit int) *Conditions {
	s.limit = limit
	s.limitValid = true
	return s
}

func (s *Conditions) Offset(offset int) *Conditions {
	s.offset = offset
	s.offsetValid = true
	return s
}

func (s *Conditions) ToSQL() (string, []interface{}) {
	sql := &bytes.Buffer{}
	var args []interface{}

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
