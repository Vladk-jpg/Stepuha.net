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
		newErrorResponse(ctx, http.StatusBadRequest, "invalid user id")
		return
	}
	var user entities.User
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id number")
		return
	}
	user, err = handl.services.GetUserById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (handl *Handler) getYourUser(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid user id")
		return
	}
	var user entities.User
	user, err = handl.services.GetYourUser(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, "user not found")
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (handl *Handler) updateUser(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "can't identify user")
		return
	}
	var input entities.UpdateUserInput
	if err = ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = handl.services.UpdateUser(id, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
