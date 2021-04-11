package controllers

import (
	"errors"
	"net/http"

	"github.com/DamnDanielV/RestServer-Go/database"
	"github.com/DamnDanielV/RestServer-Go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

type ValidateUser struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func NewUser() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

// CreateUser crea un usuario en la base de datos
// valida si se enviaron los campos correctos
// retorna un JSON con un codigo de estado 200 y el objeto creado
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *UserRepo) CreateUser(c *gin.Context) {
	var input ValidateUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Email: input.Email}
	err1 := models.CreateUser(repository.Db, &user)
	if err1 != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUsers retorna los usuarios almacenados en la base e datos
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *UserRepo) GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUser busca y retorna un usuario especificado por su id
// en caso de no encontrarse el usuario retorna un código de estado 404
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *UserRepo) GetUser(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var user models.User
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser actualiza y retorna un usuario especificado por su id
// en caso de no encontrarse el usuario retorna un código de estado 404
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *UserRepo) UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&user)
	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser elimina un usuario especificado por su id
// en caso de fallo retorna un código de estado 500 y el mensaje de error en formato JSON
func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
