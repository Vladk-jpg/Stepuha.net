package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (handl *Handler) freezeUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id number")
	}
	isModer, err := handl.services.CheckIfModer(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	if !isModer {
		newErrorResponse(ctx, http.StatusForbidden, "user hasn't got permissions")
	}

	err = handl.services.FreezeUser(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
