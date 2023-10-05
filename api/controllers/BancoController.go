package controllers

import (
	"api/api/entities"
	"api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BancoController struct{}

func NewBancoController() *BancoController {
	return &BancoController{}
}

func (bc *BancoController) CreateBanco(c *gin.Context) {
	var banco entities.Banco
	if err := c.ShouldBindJSON(&banco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateBanco(&banco); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, banco)
}

func (bc *BancoController) FindAll(c *gin.Context) {
	bancos, err := database.GetAllBancos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bancos)
}

func (bc *BancoController) FindByID(c *gin.Context) {
	bancoID := c.Param("id")
	banco, err := database.GetBanco(bancoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if banco == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banco not found"})
		return
	}

	c.JSON(http.StatusOK, banco)
}

func (bc *BancoController) UpdateBanco(c *gin.Context) {
	bancoID := c.Param("id")

	var banco entities.Banco
	if err := c.ShouldBindJSON(&banco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.UpdateBanco(bancoID, &banco); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, banco)
}

func (bc *BancoController) DeleteBanco(c *gin.Context) {
	bancoID := c.Param("id")
	if err := database.DeleteBanco(bancoID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
