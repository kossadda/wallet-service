package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kossadda/wallet-service/pkg/service"
)

// Handler is a struct that holds the service for handling HTTP requests.
type Handler struct {
	services *service.Service
}

// New creates a new handler with the given service.
// Returns: *Handler
func New(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

// InitRoutes initializes the routing for the API.
// Returns: *gin.Engine
func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	api := router.Group("/api/v1")
	{
		api.POST("/wallet", h.handleWalletOperation)
		api.GET("/wallets/:walletId", h.handleGetWalletBalance)
	}

	return router
}
