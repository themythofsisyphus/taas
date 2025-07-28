package router

import (
	"taas/handler"
	"taas/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine, services *service.Services) {

	tagHandler := handler.NewTagHandler(services.Tag)
	entityHandler := handler.NewEntityHandler(services.Entity)
	tagMappingHandler := handler.NewTagMappingHandler(services.TagMapping)
	tenantHandler := handler.NewTenantHandler(services.Tenant)

	apis := route.Group("/api")
	{
		tagAPIs := apis.Group("/tags")
		{
			tagAPIs.GET("", tagHandler.ListTags)
			tagAPIs.POST("", tagHandler.CreateTag)
			tagAPIs.GET("/:id", tagHandler.GetTagByID)
			tagAPIs.PUT("/:id", tagHandler.UpdateTag)
			tagAPIs.DELETE("/:id", tagHandler.DeleteTag)
			tagAPIs.GET("/search", tagHandler.SearchTags)
		}

		entitiesAPIs := apis.Group("/entities")
		{
			entitiesAPIs.GET("", entityHandler.ListEntities)
			entitiesAPIs.POST("", entityHandler.CreateEntity)
			entitiesAPIs.DELETE("/:type", entityHandler.DeleteEntity)
		}

		tagMappingsAPIs := apis.Group("/:entity_type/tag_mappings")
		{
			tagMappingsAPIs.GET("/:id", tagMappingHandler.ListTagMappings)
			tagMappingsAPIs.POST("/:id", tagMappingHandler.CreateTagMappings)
			tagMappingsAPIs.DELETE("/:id", tagMappingHandler.DeleteTagMappings)
		}

		tenantAPIs := apis.Group("/tenants")
		{
			tenantAPIs.POST("", tenantHandler.CreateTenant)
			tenantAPIs.GET("/:id", tenantHandler.GetTenantByID)
			tenantAPIs.DELETE("/:id", tenantHandler.DeleteTenant)
		}
	}

}
