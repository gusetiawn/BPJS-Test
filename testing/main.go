package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Transaction struct {
	ID        int     `json:"id"`
	Customer  string  `json:"customer"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Timestamp string  `json:"timestamp"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	apiURL := "http://localhost:8080/transactions"

	for i := 1; i <= 1000; i++ {
		transaction := Transaction{
			ID:        i,
			Customer:  fmt.Sprintf("Customer%d", i),
			Quantity:  rand.Intn(10) + 1,
			Price:     float64(rand.Intn(100)+1) + rand.Float64(),
			Timestamp: "2022-08-23T15:30:00Z",
		}

		transactionData, err := json.Marshal(transaction)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(transactionData))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		defer response.Body.Close()

		fmt.Printf("Response Status: %s\n", response.Status)
	}
	elapsed := time.Since(start)
	fmt.Printf("Waktu eksekusi: %s\n", elapsed)
}
