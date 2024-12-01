package handler

import (
	"fmt"
	"github.com/kossadda/wallet-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleWalletOperation handles the wallet operation request.
// Returns:
// - 200 OK: Success operation with user ID
// - 400 Bad Request: Invalid request or insufficient funds
// - 404 Not Found: No wallet with ID
// - 500 Internal Server Error: Internal server error
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

// handleGetWalletBalance handles the request to get the wallet balance.
// Returns:
// - 200 OK: Wallet balance
// - 404 Not Found: No wallet with ID
// - 500 Internal Server Error: Internal server error
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
