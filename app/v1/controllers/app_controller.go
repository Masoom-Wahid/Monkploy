package controllers

import (
	"fmt"
	"net/http"
	"platform/app/v1/repositories"
	"platform/app/v1/requests"
	"platform/pkg/validation"

	"github.com/gin-gonic/gin"
)

type AppController interface {
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type appController struct {
	repoSupplier repositories.RepoSupplier
}

func NewAppController(repo repositories.RepoSupplier) AppController {
	return &appController{
		repoSupplier: repo,
	}
}

func (controller *appController) List(ctx *gin.Context) {
	apps, err := controller.repoSupplier.AppRepository().List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": apps})
}

func (controller *appController) Create(ctx *gin.Context) {
	var credential requests.AppRequest
	ctx.ShouldBind(&credential)

	errors := validation.Validation(credential)
	if len(errors) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errors,
		})
		return
	}

	app := credential.ToAppModel()
	controller.repoSupplier.AppRepository().Create(&app)
	ctx.JSON(200, gin.H{
		"nice": "nice",
	})
}

func (controller *appController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "Missing ID parameter",
		})
		return
	}

	deleted := controller.repoSupplier.AppRepository().Delete(id)

	if !deleted {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "App not found or could not be deleted",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (service *appController) Hello(ctx *gin.Context) {
	fmt.Printf("here")
	ctx.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
