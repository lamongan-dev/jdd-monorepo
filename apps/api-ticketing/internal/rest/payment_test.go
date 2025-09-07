package rest_test

// func TestPaymentE2E(t *testing.T) {
// 	kit := NewTestKit(t)

// 	// Wire the routes and services
// 	ipaymuRepo := http_repo.NewIpaymuRepository()
// 	paymentService := service.NewPaymentService(ipaymuRepo)
// 	rest.NewPaymentHandler(kit.Echo.Group("/api/v1"), paymentService)

// 	// Now start the test server
// 	kit.Start(t)

// 	// Create
// 	createReq := domain.PaymentBody{
// 		Product:     "Test Product",
// 		Qty:         1,
// 		Price:       10000,
// 		Amount:      10000,
// 		ReferenceId: "ref-123",
// 		BuyerName:   "John Doe",
// 		BuyerEmail:  "john@example.com",
// 	}
// 	cre, code := doRequest[domain.PaymentResponse](
// 		t, http.MethodPost,
// 		kit.BaseURL+"/api/v1/payment",
// 		createReq,
// 	)

// 	require.Equal(t, http.StatusCreated, code)
// 	require.Equal(t, 200, cre.Status)
// 	payment := cre.Data
// 	require.NotEmpty(t, payment.URL)
// 	require.NotEmpty(t, payment.SessionID)
// }
