package repository

import (
	"errors"

	"github.com/fajartd02/golang-devcamp/core/entity"
	"github.com/gin-gonic/gin"
)

var ErrRecordUserNotFound = errors.New("record user not found")

type UserRepository interface {
	FindAll(c *gin.Context) ([]entity.User, error)
	Create(c *gin.Context) error
}
