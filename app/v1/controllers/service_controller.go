package controllers

import (
	"net/http"
	"platform/app/v1/repositories"
	"platform/app/v1/requests"
	"platform/pkg/validation"

	"github.com/gin-gonic/gin"
)

type ServiceController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type serviceController struct {
	repoSupplier repositories.RepoSupplier
}

func NewServiceController(repo repositories.RepoSupplier) ServiceController {
	return &serviceController{
		repoSupplier: repo,
	}
}

func (controller *serviceController) List(ctx *gin.Context) {
	appId := ctx.Param("appId")

	apps, err := controller.repoSupplier.ServiceRepository().List(appId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": apps})
}

func (controller *serviceController) Create(ctx *gin.Context) {
	appId := ctx.Param("appId")

	var credential requests.ServiceRequest
	ctx.ShouldBind(&credential)

	errors := validation.Validation(credential)
	if len(errors) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errors,
		})
		return
	}

	service := credential.ToServiceModel()
	_, err := controller.repoSupplier.ServiceRepository().Create(&service, appId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"nice": "nice",
	})
}

func (controller *serviceController) Delete(ctx *gin.Context) {
	appId := ctx.Param("appId")

	id := ctx.Param("serviceId")

	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "Missing ID parameter",
		})
		return
	}

	err := controller.repoSupplier.ServiceRepository().Delete(id, appId)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}


