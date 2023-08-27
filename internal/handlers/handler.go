package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gusetiawn/BPJS-Test/internal/database"
	"github.com/gusetiawn/BPJS-Test/internal/model"
	"github.com/gusetiawn/BPJS-Test/internal/service"
	"net/http"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler() *TransactionHandler {
	db := database.NewDatabaseConnection()
	transactionService := service.NewTransactionService(db)

	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction model.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go h.transactionService.ProcessTransaction(transaction)

	c.JSON(http.StatusAccepted, gin.H{"message": "Transaction processing started"})
}

func (h *TransactionHandler) CreateBulkTransactions(c *gin.Context) {
	var request model.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go h.transactionService.ProcessBulkTransactions(request)

	c.JSON(http.StatusAccepted, gin.H{"message": "Transaction processing started"})
}

func (h *TransactionHandler) CreateBulkTransactionsWithGoroutine(c *gin.Context) {
	var request model.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go h.transactionService.ProcessBulkTransactions(request)

	c.JSON(http.StatusAccepted, gin.H{"message": "Transaction processing started with goroutine"})
}
