package router

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	server   = "39.106.139.191:6379"
	password = "renketong"
	db       = "0"
	Pool     = newPool(server, password, db)
	//redisServer = flag.String("redisServer", "39.106.139.191:6379")
)

func newPool(server, password, db string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}
