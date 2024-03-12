package entity

import (
	"errors"
)

type OrderRequest struct {
	OrderID  string  `json:"orderId"`
	CardHash string  `json:"cardHash"`
	Total    float64 `json:"total"`
}

func NewOrderRequest(orderID, cardHash string, total float64) *OrderRequest {
	return &OrderRequest{
		OrderID:  orderID,
		CardHash: cardHash,
		Total:    total,
	}
}

func (o *OrderRequest) Validate() error {
	if o.OrderID == "" {
		return errors.New("orderId is required")
	}
	if o.CardHash == "" {
		return errors.New("cardHash is required")
	}
	if o.Total <= 0 {
		return errors.New("total must be greater than 0")
	}
	return nil
}

type OrderResponse struct {
	OrderID string `json:"orderId"`
	Status  string `json:"status"`
}

func NewOrderResponse(orderId, status string) *OrderResponse {
	return &OrderResponse{
		OrderID: orderId,
		Status:  status,
	}
}

func (o *OrderRequest) Process() (*OrderResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}
	orderResponse := NewOrderResponse(o.OrderID, "FAILED")
	if o.Total <= 1000.00 {
		orderResponse.Status = "PAID"
	}
	return orderResponse, nil
}
