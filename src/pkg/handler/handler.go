package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-service/pkg/service"
)

type Handler struct {
	services *service.Service
}

func New(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/wallet", h.handleWalletOperation)
		api.GET("/wallets/:walletId", h.handleGetWalletBalance)
	}

	return router
}
