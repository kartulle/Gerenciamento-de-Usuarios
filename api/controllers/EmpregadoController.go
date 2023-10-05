package controllers

import (
	"api/api/entities"
	"api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmpregadoController struct{}

func NewEmpregadoController() *EmpregadoController {
	return &EmpregadoController{}
}

func (ec *EmpregadoController) CreateEmpregado(c *gin.Context) {
	var empregado entities.Empregado
	if err := c.ShouldBindJSON(&empregado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateEmpregado(&empregado); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, empregado)
}

func (ec *EmpregadoController) FindAll(c *gin.Context) {
	empregados, err := database.GetAllEmpregados()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, empregados)
}

func (ec *EmpregadoController) FindByID(c *gin.Context) {
	empregadoID := c.Param("id")
	empregado, err := database.GetEmpregado(empregadoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if empregado == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empregado not found"})
		return
	}

	c.JSON(http.StatusOK, empregado)
}

func (ec *EmpregadoController) UpdateEmpregado(c *gin.Context) {
	empregadoID := c.Param("id")

	var empregado entities.Empregado
	if err := c.ShouldBindJSON(&empregado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.UpdateEmpregado(empregadoID, &empregado); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, empregado)
}

func (ec *EmpregadoController) DeleteEmpregado(c *gin.Context) {
	empregadoID := c.Param("id")
	if err := database.DeleteEmpregado(empregadoID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
