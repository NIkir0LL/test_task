package handler

import (
	"database/sql"
	"log" // Добавлен импорт для логирования
	"net/http"
	"strconv"
	"user-api/internal/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo UserRepository
}

type UserRepository interface {
	Create(user *model.User) error
	Get(id int) (*model.User, error)
	Update(user *model.User) error
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("CreateUser: Invalid input - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(&user); err != nil {
		log.Printf("CreateUser: Failed to create user - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("CreateUser: Successfully created user - ID: %d, Name: %s, Email: %s", user.ID, user.Name, user.Email)
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("GetUser: Invalid ID - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.repo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("GetUser: User not found - ID: %d", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		log.Printf("GetUser: Failed to get user - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("GetUser: Successfully retrieved user - ID: %d, Name: %s, Email: %s", user.ID, user.Name, user.Email)
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("UpdateUser: Invalid ID - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("UpdateUser: Invalid input - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id

	if err := h.repo.Update(&user); err != nil {
		log.Printf("UpdateUser: Failed to update user - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("UpdateUser: Successfully updated user - ID: %d, Name: %s, Email: %s", user.ID, user.Name, user.Email)
	c.JSON(http.StatusOK, user)
}