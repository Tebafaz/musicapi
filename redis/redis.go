package redis

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

//Redis a
var redisConn redis.Conn

//InitRedis a
func InitRedis() {
	var err error
	redisConn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to Redis created")
}

//CloseRedis a
func CloseRedis() {
	err := redisConn.Close()
	if err != nil {
		log.Fatal("redis refused to close")
	}
}
