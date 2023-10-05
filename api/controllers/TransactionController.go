package controllers

import (
	"api/api/entities"
	"api/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionController struct {
	transactions []entities.Transaction
}

func NewTransactionController() *TransactionController {
	transactionController := &TransactionController{
		transactions: []entities.Transaction{},
	}

	allTransactions, err := database.GetAllTransactions()

	if err != nil {
		// Proper error handling, e.g., log the error and decide how to handle it
		// fmt.Printf("Error fetching transactions from the database: %v\n", err)
		// You can choose to return an error or handle it in your application logic.

		// If you decide to return an error:
		// return nil, err
	} else {
		transactionController.transactions = allTransactions
	}

	return transactionController
}

func CreateTransaction(c *gin.Context, transactionController *TransactionController) {
	var transaction entities.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new Transaction instance and save it to the database
	newTransaction := entities.NewTransaction(transaction.Quantia, transaction.Descricao, transaction.SenderId, transaction.ReceiverId)

	if err := database.CreateTransaction(newTransaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transactionController.transactions = append(transactionController.transactions, *newTransaction)

	c.JSON(http.StatusCreated, newTransaction)
}

func GetTransaction(c *gin.Context) {
	transactionID := c.Param("id")
	if _, err := uuid.Parse(transactionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	// Retrieve the Transaction from the database by ID
	transaction, err := database.GetTransaction(transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
func UpdateTransaction(c *gin.Context, transactionController *TransactionController) {
	transactionID := c.Param("id")

	if _, err := uuid.Parse(transactionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction entities.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure that transaction.Sender and transaction.Receiver are valid UUIDs
	if _, err := uuid.Parse(transaction.SenderId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Sender ID"})
		return
	}
	if _, err := uuid.Parse(transaction.ReceiverId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Receiver ID"})
		return
	}

	updatedTransaction, err := database.UpdateTransaction(transactionID, &transaction)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	for i, u := range transactionController.transactions {
		if u.ID == updatedTransaction.ID {
			transactionController.transactions[i] = *updatedTransaction
			break
		}
	}

	c.JSON(http.StatusOK, updatedTransaction)
}

func DeleteTransaction(c *gin.Context, transactionController *TransactionController) {
	transactionID := c.Param("id")
	if _, err := uuid.Parse(transactionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	if err := database.DeleteTransaction(transactionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Remove the deleted user from the UserController's users slice
	for i, transaction := range transactionController.transactions {
		if transaction.ID == transactionID {
			transactionController.transactions = append(transactionController.transactions[:i], transactionController.transactions[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}

func (u *TransactionController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, u.transactions)
}
