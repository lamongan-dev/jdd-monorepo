package service

import (
	"api-ticketing/domain"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type IpaymuRepository interface {
	CreatePayment(body []byte, payload domain.PaymentHeader) (*domain.PaymentResponse, error)
}

type IPaymuService struct {
	repo IpaymuRepository
}

func NewPaymentService(repo IpaymuRepository) *IPaymuService {
	return &IPaymuService{
		repo: repo,
	}
}

func (s *IPaymuService) ProcessPayment(payment domain.PaymentBody) (*domain.PaymentResponse, error) {
	apiKey := os.Getenv("IPAYMU_API_KEY")
	va := os.Getenv("IPAYMU_VA")
	notifyUrl := os.Getenv("CALLBACK_URL")

	var url string
	if os.Getenv("APP_ENV") == "production" {
		url = "https://my.ipaymu.com/api/v2/payment"
	} else {
		url = "https://sandbox.ipaymu.com/api/v2/payment"
	}

	body := map[string]interface{}{
		"product":     []interface{}{payment.Product},
		"qty":         []interface{}{payment.Qty},
		"price":       []interface{}{payment.Price},
		"amount":      payment.Amount,
		"notifyUrl":   notifyUrl,
		"referenceId": payment.ReferenceId,
		"buyerName":   payment.BuyerName,
		"buyerEmail":  payment.BuyerEmail,
	}

	// SHA256 of body
	bodyJSON, _ := json.Marshal(body)
	bodyHash := sha256.Sum256(bodyJSON)
	bodyEncrypt := hex.EncodeToString(bodyHash[:])

	// String to sign
	stringToSign := fmt.Sprintf("POST:%s:%s:%s", va, bodyEncrypt, apiKey)

	// HMAC SHA256 signature
	hm := hmac.New(sha256.New, []byte(apiKey))
	hm.Write([]byte(stringToSign))
	signature := hex.EncodeToString(hm.Sum(nil))

	// Timestamp format: yyyyMMddHHmmss
	timestamp := time.Now().UTC().Format("20060102150405")

	// Send HTTP request
	resp, err := s.repo.CreatePayment(bodyJSON, domain.PaymentHeader{
		Va:        va,
		Signature: signature,
		Timestamp: timestamp,
		Url:       url,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}
