package routes

import (
	"platform/app/v1/controllers"
	"platform/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.RouterGroup, cs controllers.ControllerSupplier, ms middlewares.MiddlewareSupplier) {
	v1 := router.Group("/api/v1")

}
