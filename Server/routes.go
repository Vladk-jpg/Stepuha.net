package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRouts(r *gin.Engine) {
	r.POST("/register", registerHandler)
}

func registerHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Здесь вам нужно сохранить пользователя в базу данных или другое хранилище данных

	// Отправка ответа об успешной регистрации
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
