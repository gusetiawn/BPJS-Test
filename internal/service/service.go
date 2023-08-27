package service

import (
	"database/sql"
	"fmt"
	"github.com/gusetiawn/BPJS-Test/internal/model"
	"github.com/gusetiawn/BPJS-Test/internal/repository"
	"log"
	"sync"
	"time"
)

type TransactionService struct {
	repo *repository.TransactionRepository
	wg   sync.WaitGroup
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{
		repo: repository.NewTransactionRepository(db),
	}
}

func (s *TransactionService) ProcessTransaction(t model.Transaction) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		err := s.repo.InsertTransaction(t)
		if err != nil {
			log.Printf("Error inserting transaction with ID %d: %v\n", t.ID, err)
		}
	}()
}

func (s *TransactionService) ProcessBulkTransactions(req model.Request) {
	start := time.Now()

	err := s.repo.InsertBulkTransaction(req.Data)
	if err != nil {
		log.Printf("Error inserting transaction with ID %d: %v\n", req.RequestID, err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Waktu eksekusi: %s\n", elapsed)

}

func (s *TransactionService) ProcessBulkTransactionsWithGoroutine(req model.Request) {
	start := time.Now()

	err := s.repo.InsertBulkTransactionWithGoroutine(req.Data)
	if err != nil {
		log.Printf("Error inserting transaction with ID %d: %v\n", req.RequestID, err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Waktu eksekusi: %s\n", elapsed)

}
