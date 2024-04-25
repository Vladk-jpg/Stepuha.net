package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handl *Handler) createGood(ctx *gin.Context) {
	id, _ := ctx.Get(userCtx)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
