package routes

import (
	controllers "api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	userController := controllers.NewUserController()
	accountController, err := controllers.NewAccountController()
	transactionController := controllers.NewTransactionController()

	if err != nil {
		// Handle the error, e.g., log it or return an error response
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Essa é a página principal. Hello world!",
		})
	})

	v1 := router.Group("/v1")
	{
		v1.GET("/users", userController.FindAll)

		userRoutes := v1.Group("/users")
		{
			userRoutes.POST("/", func(c *gin.Context) {
				controllers.CreateUser(c, userController)
			})

			userRoutes.GET("/:id", controllers.GetUser)
			userRoutes.PUT("/:id", func(c *gin.Context) {
				controllers.UpdateUser(c, userController)
			})
			userRoutes.DELETE("/:id", func(c *gin.Context) {
				controllers.DeleteUser(c, userController)
			})

		}

		v1.GET("/transactions", transactionController.FindAll)

		transactionsRoutes := v1.Group("/transactions")
		{
			transactionsRoutes.POST("/", func(c *gin.Context) {
				controllers.CreateTransaction(c, transactionController)
			})

			transactionsRoutes.GET("/:id", controllers.GetTransaction)
			transactionsRoutes.PUT("/:id", func(c *gin.Context) {
				controllers.UpdateTransaction(c, transactionController)
			})
			transactionsRoutes.DELETE("/:id", func(c *gin.Context) {
				controllers.DeleteTransaction(c, transactionController)
			})
		}

		// Contas
		accounts := v1.Group("/accounts")
		{
			accounts.POST("/create", accountController.CreateAccount)
			accounts.GET("/:numeroConta", accountController.GetAccount)
			accounts.PUT("/:numeroConta", accountController.UpdateAccount)
			accounts.DELETE("/:numeroConta", accountController.DeleteAccount)
		}

	}

	return v1
}
