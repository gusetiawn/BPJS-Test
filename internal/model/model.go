package model

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Customer  string    `json:"customer"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}

type Request struct {
	RequestID int           `json:"request_Id"`
	Data      []Transaction `json:"data"`
}
