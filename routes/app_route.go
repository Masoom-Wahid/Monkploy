package routes

import (
	"platform/app/v1/controllers"
	"platform/middlewares"

	"github.com/gin-gonic/gin"
)

func setupAppRoutes(route *gin.RouterGroup, appController controllers.AppController, ms middlewares.MiddlewareSupplier) {
	route.GET("/apps", appController.List)
	route.POST("/apps", appController.Create)
	route.DELETE("/apps/:id", appController.Delete)
}
