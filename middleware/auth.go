package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Extract tenant_id from header
		// tenantID := c.GetHeader("X-Tenant-ID")
		// if tenantID == "" {
		//     c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing tenant_id"})
		//     return
		// }

		// Set tenant_id in Gin context
		c.Set("tenant_id", uint(1))
		c.Next()
	}
}
