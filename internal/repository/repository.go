package repository

import (
	"database/sql"
	"github.com/gusetiawn/BPJS-Test/internal/model"
	"log"
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

func (repo *TransactionRepository) InsertBulkTransaction(t []model.Transaction) error {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v\n", err)
		return err
	}

	query := `INSERT INTO transactions (customer, quantity, price, timestamp) VALUES ($1, $2, $3, $4)`
	for _, item := range t {
		_, err = tx.Exec(query, item.Customer, item.Quantity, item.Price, item.Timestamp)
		if err != nil {
			log.Printf("Error inserting transaction: %v\n", err)
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v\n", err)
		return err
	}

	return nil
}
