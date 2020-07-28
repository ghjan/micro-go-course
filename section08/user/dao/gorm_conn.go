package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/longjoy/micro-go-course/section08/user/config"
)

var (
	db *gorm.DB
)

func init() {
	var err error

	if config.DBAdapter != "" && config.DSN != "" {
		db, err = gorm.Open(config.DBAdapter, config.DSN)
	}

	if err == nil {
		db.LogMode(config.Debug)
	} else {
		panic(err)
	}
	db.SingularTable(true)

	return
}
