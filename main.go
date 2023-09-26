package main

import (
	routes "api/api/routes"
	"api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run("localhost:3001")
}
