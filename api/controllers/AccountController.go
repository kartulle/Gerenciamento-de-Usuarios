package controllers

import (
	"net/http"

	"api/api/entities"
	"api/database"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accounts map[string]*entities.Account
}

func NewAccountController() (*AccountController, error) {
	accountController := &AccountController{
		accounts: make(map[string]*entities.Account),
	}

	// Fetch all accounts from the database
	allAccounts, err := database.GetAllAccounts()
	if err != nil {
		return nil, err
	}

	// Populate the accounts map
	for _, account := range allAccounts {
		accountController.accounts[account.NumeroConta] = account
	}

	return accountController, nil
}

func (ac *AccountController) CreateAccount(c *gin.Context) {
	numeroConta := c.PostForm("numeroConta")
	saldo := c.PostForm("saldo")
	tipoConta := c.PostForm("tipoConta")

	donoObj := entities.NewUser("Test", "User", "Test address in Maceió", "82-9xxxx-xxxx", nil)

	newAccount := entities.NewAccount(numeroConta, saldo, tipoConta, donoObj, nil)

	if err := database.CreateAccount(newAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ac.accounts[numeroConta] = newAccount

	c.JSON(http.StatusCreated, newAccount)
}

func (ac *AccountController) GetAccount(c *gin.Context) {
	numeroConta := c.Param("numeroConta")

	account, err := database.GetAccount(numeroConta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if account == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}
func (ac *AccountController) UpdateAccount(c *gin.Context) {
	numeroConta := c.Param("numeroConta")
	account, found := ac.accounts[numeroConta]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	var updateData struct {
		Saldo     string `json:"saldo"`
		TipoConta string `json:"tipoConta"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// At this point, you can update the account safely
	if err := database.UpdateAccount(numeroConta, updateData.Saldo, updateData.TipoConta); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the local account
	account.Saldo = updateData.Saldo
	account.TipoConta = updateData.TipoConta

	c.JSON(http.StatusOK, account)
}

// Delete an account by NumeroConta
func (ac *AccountController) DeleteAccount(c *gin.Context) {
	numeroConta := c.Param("numeroConta")

	// Attempt to delete the account from the database
	if err := database.DeleteAccount(numeroConta); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete the account from the local map (if needed)
	delete(ac.accounts, numeroConta)

	c.JSON(http.StatusNoContent, nil)
}
