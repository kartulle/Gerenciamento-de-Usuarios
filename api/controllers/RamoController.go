package controllers

import (
	"api/api/entities"
	"api/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RamoController struct{}

func NewRamoController() *RamoController {
	return &RamoController{}
}

func (rc *RamoController) CreateRamo(c *gin.Context) {
	var ramo entities.Ramo
	if err := c.ShouldBindJSON(&ramo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ramo.ID = uuid.New().String()

	if err := database.CreateRamo(&ramo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ramo)
}

func (rc *RamoController) FindAll(c *gin.Context) {
	ramos, err := database.GetAllRamos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ramos)
}

func (rc *RamoController) FindByID(c *gin.Context) {
	ramoID := c.Param("id")
	ramo, err := database.GetRamo(ramoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if ramo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ramo not found"})
		return
	}

	c.JSON(http.StatusOK, ramo)
}

func (rc *RamoController) UpdateRamo(c *gin.Context) {
	ramoID := c.Param("id")

	var ramo entities.Ramo
	if err := c.ShouldBindJSON(&ramo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRamo, err := database.UpdateRamo(ramoID, &ramo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRamo)
}

func (rc *RamoController) DeleteRamo(c *gin.Context) {
	ramoID := c.Param("id")
	if err := database.DeleteRamo(ramoID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
