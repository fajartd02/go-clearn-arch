package repository

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/golang-devcamp/core/entity"
)

var ErrRecordProductNotFound = errors.New("record product not found")

type ProductRepository interface {
	FindAll(c *gin.Context) ([]entity.Product, error)
	FindSingle(c *gin.Context) (entity.Product, error)
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}
