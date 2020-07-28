package config

import (
	"errors"
	"fmt"
	"github.com/longjoy/micro-go-course/section08/user/utils"
	"net/url"
	"os"
	"strconv"
)

const (
	SHORTUUIDLen = 22
)

var (
	DBAdapter  string
	DBDriver   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	Debug      bool
	DSN        string
	DBMaxOpen  int64
	DBMaxIdle  int64
)

func init() {
	DBAdapter = utils.GetEnv("DB_ADAPTER", "mysql")
	DBDriver = utils.GetEnv("DB_DRIVER", DBAdapter)
	DBHost = utils.GetEnv("DB_HOST", "localhost")
	DBPort = utils.GetEnv("DB_PORT", "3306")
	DBUser = utils.GetEnv("DB_USER", "root")
	DBPassword = url.QueryEscape(utils.GetEnv("DB_PASSWORD", "123456"))
	DBName = utils.GetEnv("DB_NAME", "user")
	Debug, _ = strconv.ParseBool(utils.GetEnv("DEBUG", "false"))
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DBUser, DBPassword, DBHost, DBPort, DBName)
	if DBAdapter == "mysql" {
		DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", DBUser, DBPassword, DBHost,
			DBPort, DBName)
		// DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	} else if DBAdapter == "postgres" {
		DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBUser, DBPassword, DBHost,
			DBPort, DBName)
	} else if DBAdapter == "sqlite3" {
		DSN = fmt.Sprintf("%s/%s", os.TempDir(), DBName)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	DBMaxOpen, _ = strconv.ParseInt(utils.GetEnv("DB_MAX_OPEN", "100"), 10, 0)
	DBMaxIdle, _ = strconv.ParseInt(utils.GetEnv("DB_MAX_IDLE", "30"), 10, 0)

	if Debug {
		fmt.Printf("DSN:%s\n", DSN)
	}
}
