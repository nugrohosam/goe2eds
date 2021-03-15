package middlewares

import (
	"github.com/gin-gonic/gin"
)

// CanAccessBy using for ..
func CanAccessBy(s []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessAll := false
		if accessAll {
			c.Next()
			return
		}

		c.Abort()
	}
}
