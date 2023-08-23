package main

import (
	routes "api/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run("localhost:3001")
}
