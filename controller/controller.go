package controller

import (
	"midtrans/dto"
	"midtrans/entity"
	"midtrans/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MidtransController struct {
	midtransService entity.MidtransService
}

func NewMidtransController(midtransService entity.MidtransService) *MidtransController {
	return &MidtransController{
		midtransService: midtransService,
	}
}

func (mc *MidtransController) CreateTransaction(c *gin.Context) {
	var request dto.MidtransRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionData := dto.MapMidtransRequestToTransaction(request)

	response, err := mc.midtransService.Create(transactionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success", response))
}
