package gorm_conn

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/longjoy/micro-go-course/section08/user/config"
)

var (
	gormDB *gorm.DB
)

func Init() (gormDB *gorm.DB, err error) {

	if config.DBAdapter != "" && config.DSN != "" {
		gormDB, err = gorm.Open(config.DBAdapter, config.DSN)
	}
	gormDB.LogMode(config.Debug)
	gormDB.SingularTable(true)

	return
}

//GormDBConn:返回*gorm.DB
func GormDBConn() *gorm.DB {
	return gormDB
}
