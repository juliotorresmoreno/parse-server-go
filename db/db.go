package db

import (
	"github.com/go-xorm/xorm"
)

var _driver string
var _dsn string

func SetDefaultConf(driver, dsn string) {
	_driver = driver
	_dsn = dsn
}

func NewConnection() (*xorm.Engine, error) {
	return xorm.NewEngine(_driver, _dsn)
}
