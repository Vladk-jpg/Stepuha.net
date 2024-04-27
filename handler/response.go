package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(ctx *gin.Context, returnCode int, message string) {
	log.Printf(message)
	ctx.AbortWithStatusJSON(returnCode, errorResponse{message})
}
