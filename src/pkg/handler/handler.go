package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/kossadda/wallet-service"
)

type Handler struct {
	// insert database field
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/wallet", h.handleWalletOperation)
		api.GET("/wallets/:walletId", h.handleGetWalletBalance)
	}

	return router
}

func (h *Handler) handleWalletOperation(ctx *gin.Context) {
	var request models.WalletRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insert database wallet operation logic

	ctx.JSON(http.StatusOK, gin.H{"message": "Operation successful"})
}

func (h *Handler) handleGetWalletBalance(ctx *gin.Context) {
	walletID := ctx.Param("walletId")

	// insert database get balance sum

	ctx.JSON(http.StatusOK, gin.H{"walletId": walletID, "balance": 1000})
}
