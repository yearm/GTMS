package boot

import (
	"GTMS/conf"
	"github.com/go-redis/redis"
)

var CACHE *redis.Client

func ConnectRedis() {
	config := conf.GetRedisConfig()
	option := redis.Options{
		Network:    "tcp",
		Addr:       config.Host + ":" + config.Port,
		Password:   config.Password,
		DB:         config.Db,
		MaxRetries: 3,
		PoolSize:   32,
	}
	CACHE = redis.NewClient(&option)
}
