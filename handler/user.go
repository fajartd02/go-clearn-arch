package handler

import (
	"net/http"

	"github.com/fajartd02/golang-devcamp/core/module"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUc module.UserUseCase
}

func NewUserHandler(userUc module.UserUseCase) *UserHandler {
	return &UserHandler{
		userUc: userUc,
	}
}

func (hdl *UserHandler) Create(c *gin.Context) {
	err := hdl.userUc.CreateUser(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user succesfully created"})
}

func (hdl *UserHandler) GetAll(c *gin.Context) {
	Users, err := hdl.userUc.GetUsers(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Users})
}
