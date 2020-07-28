package config

import (
	"github.com/longjoy/micro-go-course/section08/user/common"
	"github.com/longjoy/micro-go-course/section08/user/utils"
	"strconv"
)

var (
	// RedisHost : redis host:port
	RedisHost               string
	RedisPort               string
	RedisPass               string
	RedisDB                 string
	CacheExpired            int64
	RedisPoolMaxIdle        int64
	RedisPoolMaxActive      int64
	RedisDialConnectTimeout int64
	RedisDialReadTimeout    int64
	RedisDialWriteTimeout   int64
)

func init() {
	RedisHost = utils.GetEnv("CACHE_HOST", "127.0.0.1")
	RedisPort = utils.GetEnv("CACHE_PORT", "6379")
	RedisPass = utils.GetEnv("CACHE_PASSWORD", "")

	RedisDB = utils.GetEnv("CACHE_DB_NEST", "7")
	CacheExpired, _ = strconv.ParseInt(utils.GetEnv("CACHE_EXPIRED", strconv.Itoa(common.CACHE_EXPIRED)), 10, 0)
	RedisPoolMaxIdle, _ = strconv.ParseInt(utils.GetEnv("REDIS_POOL_MAX_IDLE", "1000"), 10, 0)
	RedisPoolMaxActive, _ = strconv.ParseInt(utils.GetEnv("REDIS_POOL_MAX_ACTIVE", "300"), 10, 0)
	RedisDialConnectTimeout, _ = strconv.ParseInt(utils.GetEnv("REDIS_DIAL_CONNECT_TIMEOUT", "8000"), 10, 0)
	RedisDialReadTimeout, _ = strconv.ParseInt(utils.GetEnv("REDIS_DIAL_READ_TIMEOUT", "9000"), 10, 0)
	RedisDialWriteTimeout, _ = strconv.ParseInt(utils.GetEnv("REDIS_DIAL_WRITE_TIMEOUT", "12000"), 10, 0)
}
