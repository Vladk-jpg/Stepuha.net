package handler

import (
	"Stepuha.net/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (handl *Handler) createGood(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}
	var input entities.Good
	if err := ctx.ShouldBindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handl.services.Create(userId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllGoodsResponse struct {
	Data []entities.Good `json:"data"`
}

func (handl *Handler) getAllGoods(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	goods, err := handl.services.GetAll(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, getAllGoodsResponse{
		Data: goods,
	})
}

func (handl *Handler) getGoodById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id number")
		return
	}

	good, err := handl.services.GetGoodById(userId, id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, good)
}

func (handl *Handler) deleteGood(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id number")
	}

	err = handl.services.Good.Delete(userId, id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (handl *Handler) updateGood(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id number")
	}

	var input entities.UpdateGoodInput
	if err = ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err = handl.services.Update(userId, id, input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
