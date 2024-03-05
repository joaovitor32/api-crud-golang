package controller

import (
	"api-crud-golang/internal/modules/user/infra/orm/gorm/repository"
	"api-crud-golang/internal/modules/user/services"

	"github.com/gin-gonic/gin"
)

type Controllers interface {
	SaveArchive(c *gin.Context)
}

type ControllersStruct struct {
	Services services.Services
	Repository repository.UserRepository
}


func (ctrl *ControllersStruct) SaveUser(c *gin.Context) {
	body := SaveArchiveImplementation(c)
	ctrl.Services.SaveUser(c, body, ctrl.Repository)
}
