package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//AccessLog a
func AccessLog(all bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		startedAt := time.Now()
		c.Next()

		elapsed := time.Now().Sub(startedAt)

		statusCode := c.Writer.Status()

		if all || statusCode >= 400 {
			fmt.Println(startedAt.Format("2006-01-02 15:04:05.99999999"), url, statusCode, elapsed.String())
		}
	}
}
