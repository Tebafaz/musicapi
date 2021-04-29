package middlewares

import (
	"errors"
	"musicapi/redis"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//CheckToken A
func CheckToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithError(http.StatusBadRequest, errors.New("bad authorization schema, need Bearer"))
		return
	}

	if len(authHeader) < 9 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("token not passed"))
		return
	}

	token := authHeader[7:]
	username, err := redis.GetUsername(token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if username == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("username", username)
}
