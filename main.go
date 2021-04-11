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

// setupRouter configura las rutas a sus respectivos controladores
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.NewUser()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	ticketRepo := controllers.NewTicket()
	r.POST("/tickets", ticketRepo.CreateTicket)
	r.GET("/tickets", ticketRepo.GetTickets)
	r.GET("/tickets/:id", ticketRepo.GetTicket)
	r.PUT("/tickets/:id", ticketRepo.UpdateTicket)
	r.DELETE("/tickets/:id", ticketRepo.DeleteTicket)

	return r
}
