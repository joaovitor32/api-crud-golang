package services

import (
	"api-crud-golang/internal/modules/user/infra/orm/gorm/repository"
	"api-crud-golang/internal/shared/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveUserImplementation(c *gin.Context, body types.RequestBody, repository repository.UserRepositoryInterface) {

	usr, result := repository.AddUser(c, body)

	c.JSON(http.StatusOK, gin.H{
		"rowsAffected": result.RowsAffected,
		"_id":          usr.ID,
	})

}
