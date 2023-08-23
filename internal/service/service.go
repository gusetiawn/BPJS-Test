package service

import (
	"database/sql"
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

		// Simulate processing time
		time.Sleep(10 * time.Millisecond)

		err := s.repo.InsertTransaction(t)
		if err != nil {
			log.Printf("Error inserting transaction with ID %d: %v\n", t.ID, err)
		}
	}()
}

func (s *TransactionService) WaitForCompletion() {
	s.wg.Wait()
}
