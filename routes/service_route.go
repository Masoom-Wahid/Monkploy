package routes

import (
	"platform/app/v1/controllers"
	"platform/middlewares"

	"github.com/gin-gonic/gin"
)

func setupServiceRoutes(route *gin.RouterGroup, serviceController controllers.ServiceController, ms middlewares.MiddlewareSupplier) {
	services := route.Group("services")

	services.GET("/", serviceController.List)
	services.POST("/", serviceController.Create)
	services.DELETE("/:serviceId", serviceController.Delete)
}
