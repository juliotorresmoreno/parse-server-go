package acl

import "github.com/casbin/casbin"
import "github.com/casbin/casbin/model"

func NewACL() model.Model {
	m := casbin.NewModel()
	m.AddPolicy("", "", []string{})
	return m
}
