package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, return_code int, message string) {
	log.Printf(message)
	ctx.AbortWithStatusJSON(return_code, errorResponse{message})
}
