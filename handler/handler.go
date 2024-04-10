package handler

import (
	"Stepuha.net/entities"
	"Stepuha.net/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Nah, not now. Or, perhaps?..

func (handl *Handler) RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handl.signUP)
	}
}

func registerHandler(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Здесь вам нужно сохранить пользователя в базу данных или другое хранилище данных

	// Отправка ответа об успешной регистрации
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
