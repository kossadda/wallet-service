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

	id, err := h.services.Operators.BalanceChange(*request)
	if err != nil {
		switch err.Error() {
		case "no wallet with ID":
			newErrorResponse(ctx, http.StatusNotFound, err.Error())
		case "insufficient funds":
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		default:
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Success operation with %s user", id))
}

func (h *Handler) handleGetWalletBalance(ctx *gin.Context) {
	walletID := ctx.Param("walletId")

	balance, err := h.services.Operators.BalanceCheck(walletID)
	if err != nil {
		switch err.Error() {
		case "no wallet with ID":
			newErrorResponse(ctx, http.StatusNotFound, err.Error())
		default:
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"walletId": walletID, "balance": balance})
}
