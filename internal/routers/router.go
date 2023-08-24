package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gusetiawn/BPJS-Test/internal/handlers"
)

func InitRouter(router *gin.Engine) {
	// Initialize handlers
	transactionHandler := handlers.NewTransactionHandler()

	// Set up routes
	router.POST("/transactions", transactionHandler.CreateTransaction)
	router.POST("/bulkTransactions", transactionHandler.CreateBulkTransactions)
}
