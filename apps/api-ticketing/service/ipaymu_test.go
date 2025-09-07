package service_test

import (
	"api-ticketing/service"
	"api-ticketing/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaymentService(t *testing.T) {
	mockRepo := new(mocks.IpaymuRepository)
	service := service.NewPaymentService(mockRepo)

	assert.NotNil(t, service)
}
