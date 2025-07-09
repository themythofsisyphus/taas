package middleware

import (
	"context"

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
		tenantID := uint(1)
		c.Set("tenant_id", tenantID)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "tenant_id", tenantID))
		c.Next()
	}
}
