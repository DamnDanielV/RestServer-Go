package main

import (
	"net/http"

	"github.com/DamnDanielV/RestServer-Go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	ticketRepo := controllers.NewTicket()
	r.POST("/tickets", ticketRepo.CreateTicket)
	r.GET("/tickets", ticketRepo.GetTickets)
	r.GET("/tickets/:id", ticketRepo.GetTicket)
	r.PUT("/tickets/:id", ticketRepo.UpdateTicket)
	r.DELETE("/tickets/:id", ticketRepo.DeleteTicket)

	return r
}
