package service

import (
	"database/sql"
	"github.com/gusetiawn/BPJS-Test/internal/model"
	"github.com/gusetiawn/BPJS-Test/internal/repository"
	"log"
	"sync"
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

func (s *TransactionService) ProcessTransactions(t []model.Transaction) {
	s.wg.Add(len(t))

	for _, transaction := range t {
		err := s.repo.InsertTransaction(transaction)
		if err != nil {
			log.Printf("Error inserting transaction with ID %d: %v\n", transaction.ID, err)
		}
	}

}
