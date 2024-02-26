package route

import (
	"midtrans/controller"
	"midtrans/repository"
	"midtrans/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(g *gin.Engine, db *gorm.DB) {
	midtransRepository := repository.NewMidtransRepository(db)
	midtransService := service.NewMidtransService(midtransRepository)
	midtransController := controller.NewMidtransController(midtransService)

	// Tambahkan rute untuk CreateTransaction
	g.POST("/midtrans/create", midtransController.CreateTransaction)
}
