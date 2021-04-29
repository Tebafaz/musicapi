package redis

import (
	"fmt"
	"musicapi/models"

	"github.com/gomodule/redigo/redis"
)

//AddToken a
func AddToken(user models.User, token string, expiration int) error {
	_, err := redisConn.Do("SETEX", token, fmt.Sprint(expiration), user.Username)
	if err != nil {
		return err
	}
	return nil
}

//GetToken a
func GetUsername(token string) (string, error) {
	username, err := redis.Bytes(redisConn.Do("GET", token))
	return string(username), err
}
