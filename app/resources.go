package app

import (
	resources "gin-app/util"
	"fmt"
)

var (
	RedisPool resources.RedisCache
	Config resources.Config
)

const ConfigPath = "web.conf"

func init() {
	fmt.Println("load config ...")
	Config = resources.NewConfig(ConfigPath)
	fmt.Println(Config)

	fmt.Println("get redis pool...")
	RedisPool = resources.NewRedisCache(Config.GetString("redis", "redis"))
	fmt.Println(RedisPool)
}