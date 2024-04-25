package handler

import (
	"Stepuha.net/entities"
	"Stepuha.net/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func (handl *Handler) CreateGood() {
	fmt.Println("It's a scam all along")
}
func (handl *Handler) RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handl.signUp)
		auth.POST("/sign-in", handl.signIn)
	}
	api := router.Group("/api", handl.userIdentity)
	{
		userGoods := api.Group("/goods")
		{
			userGoods.POST("/", handl.createGood)
		}
	}
}

func registerHandler(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Здесь вам нужно сохранить пользователя в базу данных или другое хранилище данных
	// Nah, not now. Or, perhaps?..
	// Отправка ответа об успешной регистрации
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
