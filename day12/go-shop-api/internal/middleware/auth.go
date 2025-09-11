package middleware

import (
	"net/http"
	"strings"

	"go-shop-api/internal/auth/jwt"

	"github.com/gin-gonic/gin"
)

const ContextUserKey = "currentUser"

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}
		tokenStr := parts[1]
		claims, err := jwt.ParseAndVerify(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}
		c.Set(ContextUserKey, claims)
		c.Next()
	}
}

func GetClaims(c *gin.Context) (*jwt.Claims, bool) {
	v, ok := c.Get(ContextUserKey)
	if !ok {
		return nil, false
	}
	claims, ok := v.(*jwt.Claims)
	return claims, ok
}
