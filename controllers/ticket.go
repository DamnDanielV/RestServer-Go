package controllers

import (
	"errors"
	"net/http"

	"github.com/DamnDanielV/RestServer-Go/database"
	"github.com/DamnDanielV/RestServer-Go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TicketRepo struct {
	Db *gorm.DB
}

type ValidateTicket struct {
	User   string `json:"user" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func NewTicket() *TicketRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Ticket{})
	return &TicketRepo{Db: db}
}

// CreateTicket crea un ticket en la base de datos
// valida si se enviaron los campos correctos
// retorna un JSON con un codigo de estado 200 y el objeto creado
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *TicketRepo) CreateTicket(c *gin.Context) {
	var input ValidateTicket
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ticket := models.Ticket{User: input.User, Status: input.Status}
	err1 := models.CreateTicket(repository.Db, &ticket)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
	}
	c.JSON(http.StatusOK, ticket)
}

// GetTickets retorna los tickets almacenados en la base e datos
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *TicketRepo) GetTickets(c *gin.Context) {
	var ticket []models.Ticket
	err := models.GetTickets(repository.Db, &ticket)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// GetTicket busca y retorna un ticket especificado por su id
// en caso de no encontrarse el ticket retorna un código de estado 404
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *TicketRepo) GetTicket(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var ticket models.Ticket
	err := models.GetTicket(repository.Db, &ticket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// UpdateTicket actualiza y retorna un ticket especificado por su id
// en caso de no encontrarse el ticket retorna un código de estado 404
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *TicketRepo) UpdateTicket(c *gin.Context) {
	var ticket models.Ticket
	id, _ := c.Params.Get("id")
	err := models.GetTicket(repository.Db, &ticket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&ticket)
	err = models.UpdateTicket(repository.Db, &ticket)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// DeleteTicket elimina un ticket especificado por su id
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *TicketRepo) DeleteTicket(c *gin.Context) {
	var ticket models.Ticket
	id, _ := c.Params.Get("id")
	err := models.GetTicket(repository.Db, &ticket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteTicket(repository.Db, &ticket, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}
