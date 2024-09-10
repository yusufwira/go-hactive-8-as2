package dto

import (
	"assigment2/service/module/request_order/entity"
	"time"
)

type CreateItemRequest struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type CreateRequestOrderRequest struct {
	OrderedAt    time.Time           `json:"orderedAt"`
	CustomerName string              `json:"customerName"`
	Items        []CreateItemRequest `json:"items"`
}

type Item struct {
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Id          int       `json:"id"`
	OrderId     int       `json:"orderid"`
}

type RequestOrder struct {
	Id           int            `json:"id"`
	CustomerName string         `json:"customer_name"`
	Items        []entity.Items `json:"items"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}
