package routes

import (
	"platform/app/v1/controllers"
	"platform/middlewares"

	"github.com/gin-gonic/gin"
)

func setupAppRoutes(route *gin.RouterGroup, appController controllers.AppController, ms middlewares.MiddlewareSupplier) {
	apps := route.Group("/apps")

	apps.GET("/", appController.List)
	apps.POST("/", appController.Create)
	apps.DELETE("/:appId", appController.Delete)
}