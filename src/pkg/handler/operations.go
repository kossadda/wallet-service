package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/kossadda/wallet-service"
)

func (h *handler) handleWalletOperation(ctx *gin.Context) {
	request := models.NewRequest()

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insert database wallet operation logic

	ctx.JSON(http.StatusOK, gin.H{"message": "Operation successful"})
}

func (h *handler) handleGetWalletBalance(ctx *gin.Context) {
	walletID := ctx.Param("walletId")

	// insert database get balance sum

	ctx.JSON(http.StatusOK, gin.H{"walletId": walletID, "balance": 1000})
}
