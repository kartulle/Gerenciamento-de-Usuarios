package controllers

import (
	"api/api/entities"
	"net/http"
	"github.com/gin-gonic/gin"
)

type userController struct {
	users []entities.User
}

func NewUserController() *userController {
	return &userController{}
}

func (u *userController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, u.users)
}