package routes

import (
	controllers "api/api/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	userController := controllers.NewUserController()
	accountController, err := controllers.NewAccountController()
	transactionController := controllers.NewTransactionController()
	bancoController := controllers.NewBancoController()
	empregadoController := controllers.NewEmpregadoController()
	ramoController := controllers.NewRamoController()

	if err != nil {
		fmt.Println("Error initializing account controller")
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

		accounts := v1.Group("/accounts")
		{
			accounts.POST("/create", accountController.CreateAccount)
			accounts.GET("/:numeroConta", accountController.GetAccount)
			accounts.PUT("/:numeroConta", accountController.UpdateAccount)
			accounts.DELETE("/:numeroConta", accountController.DeleteAccount)
		}

		bancos := v1.Group("/bancos")
		{
			bancos.POST("/", bancoController.CreateBanco)
			bancos.GET("/", bancoController.FindAll)
			bancos.GET("/:id", bancoController.FindByID)
			bancos.PUT("/:id", bancoController.UpdateBanco)
			bancos.DELETE("/:id", bancoController.DeleteBanco)
		}

		empregados := v1.Group("/empregados")
		{
			empregados.POST("/", empregadoController.CreateEmpregado)
			empregados.GET("/", empregadoController.FindAll)
			empregados.GET("/:id", empregadoController.FindByID)
			empregados.PUT("/:id", empregadoController.UpdateEmpregado)
			empregados.DELETE("/:id", empregadoController.DeleteEmpregado)
		}

		ramos := v1.Group("/ramos")
		{
			ramos.POST("/", ramoController.CreateRamo)
			ramos.GET("/", ramoController.FindAll)
			ramos.GET("/:id", ramoController.FindByID)
			ramos.PUT("/:id", ramoController.UpdateRamo)
			ramos.DELETE("/:id", ramoController.DeleteRamo)
		}
	}

	return v1
}
