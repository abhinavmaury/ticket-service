package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		finish := make(chan struct{}, 1)
		defer close(finish)
		go func() {
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-time.After(t):
		case <-finish:
		}
	}
}
