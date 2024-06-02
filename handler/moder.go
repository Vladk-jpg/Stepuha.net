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

func (handl *Handler) unfreezeUser(ctx *gin.Context) {
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

	err = handl.services.UnfreezeUser(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (handl *Handler) deleteUsersGood(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid user id number")
	}

	goodId, err := strconv.Atoi(ctx.Param("goodId"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid good id number")
	}

	err = handl.services.Delete(userId, goodId)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
