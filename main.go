package main

import (
	"midtrans/database"
	"midtrans/route"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	db := database.InitPostgresDB()

	route.SetupRoutes(g, db)
	g.Run(":8080")
}
