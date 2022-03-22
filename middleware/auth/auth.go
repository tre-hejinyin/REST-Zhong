package auth

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"server/middleware/cache"
	"server/middleware/constants"
	"server/middleware/jwt"
)

// JWTAuthMiddleware
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		// get id from redis
		token, err := cache.Get(strconv.Itoa(mc.MyID))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
		if token != parts[1] {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		// set account in context. use `c.Get(constants.ID)` to obtain
		c.Set(constants.ID, mc.MyID)
		c.Next()
	}
}
