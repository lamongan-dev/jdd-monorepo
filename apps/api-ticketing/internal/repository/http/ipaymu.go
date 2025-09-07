package http

import (
	"api-ticketing/domain"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type IpaymuRepository struct {
	client *http.Client
}

func NewIpaymuRepository() *IpaymuRepository {
	return &IpaymuRepository{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (i *IpaymuRepository) CreatePayment(body []byte, payload domain.PaymentHeader) (*domain.PaymentResponse, error) {
	req, err := http.NewRequest("POST", payload.Url, bytes.NewBuffer(body))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("va", payload.Va)
	req.Header.Set("signature", payload.Signature)
	req.Header.Set("timestamp", payload.Timestamp)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, domain.ErrInternalServerError
	}

	respBytes, _ := ioutil.ReadAll(resp.Body)
	var result domain.PaymentResponse
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return nil, domain.ErrInternalServerError
	}

	return &result, nil
}
