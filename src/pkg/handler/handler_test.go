package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/service"
	mockservice "github.com/kossadda/wallet-service/pkg/service/mocks"

	"github.com/stretchr/testify/assert"
)

func TestHandleWalletOperation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOperators := mockservice.NewMockOperators(ctrl)

	services := &service.Service{Operators: mockOperators}

	h := New(services)
	router := gin.New()
	router.POST("/api/v1/wallet", h.handleWalletOperation)

	t.Run("Success", func(t *testing.T) {
		mockOperators.EXPECT().BalanceChange(gomock.Any()).Return("wallet123", nil)

		body := models.Request{
			WalletID:      "wallet123",
			OperationType: "DEPOSIT",
			Amount:        100.0,
		}
		bodyBytes, _ := json.Marshal(body)

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/wallet", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Success operation with wallet123")
	})

	t.Run("Invalid Request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/wallet", bytes.NewBuffer([]byte("invalid")))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "invalid character")
	})

	t.Run("Service Error", func(t *testing.T) {
		mockOperators.EXPECT().BalanceChange(gomock.Any()).Return("", errors.New("service error"))

		body := models.Request{
			WalletID:      "wallet123",
			OperationType: "WITHDRAW",
			Amount:        50.0,
		}
		bodyBytes, _ := json.Marshal(body)

		req, _ := http.NewRequest(http.MethodPost, "/api/v1/wallet", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "service error")
	})
}

func TestHandleGetWalletBalance(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOperators := mockservice.NewMockOperators(ctrl)

	services := &service.Service{Operators: mockOperators}

	h := New(services)
	router := gin.New()
	router.GET("/api/v1/wallets/:walletId", h.handleGetWalletBalance)

	t.Run("Success", func(t *testing.T) {
		mockOperators.EXPECT().BalanceCheck("wallet123").Return(100.0, nil)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/wallets/wallet123", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"walletId":"wallet123"`)
		assert.Contains(t, rec.Body.String(), `"balance":100`)
	})

	t.Run("Service Error", func(t *testing.T) {
		mockOperators.EXPECT().BalanceCheck("wallet123").Return(0.0, errors.New("service error"))

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/wallets/wallet123", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "service error")
	})
}
