package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenValidator(token, header string) gin.HandlerFunc {
	return func(c *gin.Context) {
		values := c.Request.Header.Values(header)
		if len(values) == 0 {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("missing token in %s header", header))
			return
		}
		if len(values) > 1 {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("cannot have more than one value for %s header", header))
			return
		}
		value := values[0]
		fmt.Println(value, token)
		if value != token {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid API key in header %s", header))
			return
		}
		c.Next()
	}
}
