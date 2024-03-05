package services

import (
	"api-crud-golang/internal/modules/user/infra/orm/gorm/repository"
	"api-crud-golang/internal/shared/types"

	"github.com/gin-gonic/gin"
)

type Services interface {
	SaveUser(c *gin.Context, body types.RequestBody, repository repository.UserRepositoryInterface)
}

func NewServices() Services {
	return &ServicesStruct{}
}

type ServicesStruct struct{}

func (s *ServicesStruct) SaveUser(c *gin.Context, body types.RequestBody, repository repository.UserRepositoryInterface) {
	SaveUserImplementation(c, body, repository)
}
