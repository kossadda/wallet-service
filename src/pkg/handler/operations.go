package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/kossadda/wallet-service"
)

func (h *Handler) handleWalletOperation(ctx *gin.Context) {
	request := models.NewRequest()

	if err := ctx.ShouldBindJSON(request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Operators.Deposit(*request)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Successfuly added %d user", id))
}

func (h *Handler) handleGetWalletBalance(ctx *gin.Context) {
	walletID := ctx.Param("walletId")

	// insert database get balance sum

	ctx.JSON(http.StatusOK, gin.H{"walletId": walletID, "balance": 1000})
}
