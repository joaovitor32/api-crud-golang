package repository

import (
	"api-crud-golang/internal/modules/user/infra/orm/gorm/database"
	userModel "api-crud-golang/internal/modules/user/infra/orm/gorm/model"

	"api-crud-golang/internal/shared/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	AddUser(c *gin.Context, user types.RequestBody) (userModel.User, *gorm.DB)
}

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (userRepo UserRepository) AddUser(c *gin.Context, user types.RequestBody) (userModel.User, *gorm.DB) {
	db, err := database.NewDatabase()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return userModel.User{}, nil
	}

	usr := userModel.User{Name: user.Name, Age: user.Age}

	result := db.Create(&usr)
	return usr, result

}
