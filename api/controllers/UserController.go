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
func (u *userController) Create(ctx *gin.Context) {
	user := entities.NewUser()

	if erro := ctx.BindJSON(&user); erro != nil {
		return
	}

	u.users = append(u.users, *user)

	ctx.JSON(http.StatusOK, user)
}
func (u *userController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, user := range u.users {
		if user.ID == id {
			u.users = append(u.users[0:index], u.users[index+1:]...)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}