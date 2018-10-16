package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/hetong12345/learn/single"
	"io"
	"net/http"
	"time"
)

var (
	server   = "39.106.139.191:6379"
	password = "renketong"
	db       = "0"
	pool     = newPool(server, password, db)
	//redisServer = flag.String("redisServer", "39.106.139.191:6379")
)

func firstPage(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "<h1>Hello World</h1>")
	conn := pool.Get()
	defer conn.Close()
	_, err := redis.Int64(conn.Do("incr", "view"))
	if err != nil {
		panic(err)
	}
	reply, _ := redis.String(conn.Do("get", "view"))
	io.WriteString(w, "您是这个页面的第"+reply+"个访问者")
}

func main() {

	//http.HandleFunc("/", firstPage)
	//http.ListenAndServe(":8000", nil)
	instance := single.GetInstance()
	instance2 := single.GetInstance()

	fmt.Println(instance)
	fmt.Println(instance2)
	//conn.Close()

}
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
