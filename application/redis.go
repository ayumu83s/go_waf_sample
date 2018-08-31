package application

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func SetupRedis(config *Config) redis.Conn {
	c, err := redis.Dial("tcp", fmt.Sprintf(":%d", config.Redis.Port))
	if err != nil {
		panic(err)
	}
	return c
}
