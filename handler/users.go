package handler

import (
	"Stepuha.net/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (handl *Handler) getUserById(ctx *gin.Context) {
	_, err := getUserId(ctx)
	if err != nil {
		return
	}
	var user entities.User
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id number")
		return
	}
	user, err = handl.services.GetUserById(id)

	ctx.JSON(http.StatusOK, user)
}
