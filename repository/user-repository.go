package repository

import (
	"errors"

	"github.com/fajartd02/golang-devcamp/core/entity"
	repository_intf "github.com/fajartd02/golang-devcamp/core/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type repositoryUser struct {
}

func NewUser() repository_intf.UserRepository {
	return &repositoryUser{}
}

func (r *repositoryUser) Create(c *gin.Context) error {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return errors.New("failed to create user")
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	User := entity.User{
		FullName: input.FullName,
		Age:      input.Age,
	}

	if err := db.Create(&User).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (r *repositoryUser) FindAll(c *gin.Context) ([]entity.User, error) {
	var users []entity.User

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	err := db.Model(&entity.User{}).Preload("Comments").Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordProductNotFound
		}
		return nil, err
	}

	return users, nil
}
