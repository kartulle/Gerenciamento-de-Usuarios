package routes

import (
	controllers "api/api/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	userController := controllers.NewUserController()

	v1:= router.Group("/v1")
	{
		v1.GET("/users", userController.FindAll)
	}

	return v1
}