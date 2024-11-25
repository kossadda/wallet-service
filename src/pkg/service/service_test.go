package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/repository"
	mockrepository "github.com/kossadda/wallet-service/pkg/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBalanceChange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOperation := mockrepository.NewMockOperation(ctrl)
	mockRepo := repository.Repository{Operation: mockOperation}
	operationService := NewOperationService(mockRepo)

	req := models.Request{
		// Initialize with appropriate values
	}

	expectedResult := "success"
	mockOperation.EXPECT().BalanceChange(req).Return(expectedResult, nil)

	result, err := operationService.BalanceChange(req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestBalanceChangeError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOperation := mockrepository.NewMockOperation(ctrl)
	mockRepo := repository.Repository{Operation: mockOperation}
	operationService := NewOperationService(mockRepo)

	req := models.Request{
		// Initialize with appropriate values
	}

	expectedError := errors.New("some error")
	mockOperation.EXPECT().BalanceChange(req).Return("", expectedError)

	_, err := operationService.BalanceChange(req)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}

func TestBalanceCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOperation := mockrepository.NewMockOperation(ctrl)
	mockRepo := repository.Repository{Operation: mockOperation}
	operationService := NewOperationService(mockRepo)

	id := "12345"
	expectedBalance := 100.0
	mockOperation.EXPECT().BalanceCheck(id).Return(expectedBalance, nil)

	balance, err := operationService.BalanceCheck(id)
	assert.NoError(t, err)
	assert.Equal(t, expectedBalance, balance)
}

func TestBalanceCheckError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOperation := mockrepository.NewMockOperation(ctrl)
	mockRepo := repository.Repository{Operation: mockOperation}
	operationService := NewOperationService(mockRepo)

	id := "12345"
	expectedError := errors.New("some error")
	mockOperation.EXPECT().BalanceCheck(id).Return(0.0, expectedError)

	_, err := operationService.BalanceCheck(id)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
