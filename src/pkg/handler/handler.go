package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-service/pkg/service"
)

func New(serv *service.Service) *handler {
	return &handler{services: serv}
}

type handler struct {
	services *service.Service
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/wallet", handleWalletOperation)
		api.GET("/wallets/:walletId", handleGetWalletBalance)
	}

	return router
}
