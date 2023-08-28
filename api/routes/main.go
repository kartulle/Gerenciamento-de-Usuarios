package routes

import (
	controllers "api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	userController := controllers.NewUserController()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Essa é a página principal. Hello world!",
		})
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/users", func(u *gin.Context) {
			u.JSON(200, gin.H{
				"message": "Essa é a página para pegar todos os usuários e retornar via JSON",
			})
		}, userController.FindAll)
		v1.POST("/user", userController.Create)
		v1.DELETE("/user/:id", userController.Delete)
	}

	return v1
}
