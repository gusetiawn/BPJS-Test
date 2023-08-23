package repository

import (
	"database/sql"
	"github.com/gusetiawn/BPJS-Test/internal/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) InsertTransaction(t model.Transaction) error {
	query := `
		INSERT INTO transactions (id, customer, quantity, price, timestamp)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := repo.db.Exec(query, t.ID, t.Customer, t.Quantity, t.Price, t.Timestamp)
	return err
}
