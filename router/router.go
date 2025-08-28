// Package router initializes and registers API routes for the application.
package router

import (
	"taas/handler"
	"taas/service"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API routes with the given Gin engine.
func RegisterRoutes(router *gin.Engine, services *service.Services) {
	tagHandler := handler.NewTagHandler(services.Tag)
	entityHandler := handler.NewEntityHandler(services.Entity)
	tagMappingHandler := handler.NewTagMappingHandler(services.TagMapping)
	tenantHandler := handler.NewTenantHandler(services.Tenant)

	api := router.Group("/api")
	{
		// Tag APIs
		tags := api.Group("/tags")
		{
			tags.GET("", tagHandler.ListTags)
			tags.POST("", tagHandler.CreateTag)
			tags.GET("/:id", tagHandler.GetTagByID)
			tags.PUT("/:id", tagHandler.UpdateTag)
			tags.DELETE("/:id", tagHandler.DeleteTag)
			tags.GET("/search", tagHandler.SearchTags)
		}

		// Entity APIs
		entities := api.Group("/entities")
		{
			entities.GET("", entityHandler.ListEntities)
			entities.POST("", entityHandler.CreateEntity)
			entities.DELETE("/:type", entityHandler.DeleteEntity)
		}

		// Tag Mapping APIs (per entity type and ID)
		tagMappings := api.Group("/:entity_type/tag_mappings")
		{
			tagMappings.GET("/:id", tagMappingHandler.ListTagMappings)
			tagMappings.POST("/:id", tagMappingHandler.CreateTagMappings)
			tagMappings.DELETE("/:id", tagMappingHandler.DeleteTagMappings)
		}

		// Tenant APIs
		tenants := api.Group("/tenants")
		{
			tenants.POST("", tenantHandler.CreateTenant)
			tenants.GET("/:id", tenantHandler.GetTenantByID)
			tenants.DELETE("/:id", tenantHandler.DeleteTenant)
		}

		// Analytics APIs
		analytics := api.Group("/analytics")
		{
			analyticsHandler := handler.NewAnalyticsHandler()
			analytics.GET("/stats", analyticsHandler.GetStats)
			analytics.GET("/endpoints", analyticsHandler.GetEndpointMetrics)
			analytics.GET("/activity", analyticsHandler.GetRecentActivity)
			analytics.GET("/traffic", analyticsHandler.GetTrafficData)
		}
	}
}
