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
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Essa é a página para pegar todos os usuários e retornar via JSON",
			})
		}, userController.FindAll)
	}

	return v1
}
