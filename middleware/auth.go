package middleware

import (
	"context"
	"net/http"
	"strings"
	"taas/config"
	"taas/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const tenantIDKey contextKey = "tenant_id"

func AuthMiddleware(config *config.JWTSecretConfig, tenantService *service.TenantService) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenMalformed
			}
			return []byte(config.Key), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		tenantIDFloat, ok := claims["tenant_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "tenant_id missing in token"})
			c.Abort()
			return
		}
		tenantID := uint(tenantIDFloat)
		_, err = tenantService.GetTenantByID(tenantID)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "tenant_id not found"})
			c.Abort()
			return
		}

		c.Set(string(tenantIDKey), tenantID)
		ctx := context.WithValue(c.Request.Context(), tenantIDKey, tenantID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
