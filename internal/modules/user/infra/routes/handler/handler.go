package pdfHandler

import (
	controler "api-crud-golang/internal/modules/user/infra/controller"
	"api-crud-golang/internal/modules/user/infra/orm/gorm/repository"
	services "api-crud-golang/internal/modules/user/services"

	//	"api-crud-golang/internal/shared/infra/handler"
	//	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandler(services services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

type Handler struct {
	services services.Services
	repository repository.UserRepository
}

func (h *Handler) saveUser(c *gin.Context) {
	controller := controler.ControllersStruct{Services: h.services, Repository: h.repository}
	controller.SaveUser(c)
}

func RoutesV1(h *Handler, router *gin.Engine) {
	pdfGroup := router.Group("user")

	pdfGroup.Use()
	{
		pdfGroup.POST("/", h.saveUser)
	}
}
