package module

import (
	"errors"
	"fmt"

	"github.com/fajartd02/golang-devcamp/core/entity"
	"github.com/fajartd02/golang-devcamp/core/repository"
	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	CreateUser(c *gin.Context) error
	GetUsers(c *gin.Context) ([]entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

var ErrUserNotFound = errors.New("user error: ")

func NewUserUsecase(repo repository.UserRepository) UserUseCase {
	return &userUsecase{repo}
}

func (em *userUsecase) CreateUser(c *gin.Context) error {
	err := em.userRepo.Create(c)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}

func (em *userUsecase) GetUsers(c *gin.Context) ([]entity.User, error) {
	data, err := em.userRepo.FindAll(c)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}

	return data, nil
}
