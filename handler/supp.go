package handler

import (
	"Stepuha.net/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getRandomGoodsResponse struct {
	Data []entities.Good `json:"data"`
}

func (handl *Handler) getRandomGoods(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "could not get user id")
		return
	}

	goods, err := handl.services.GetRandomGoods(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, getRandomGoodsResponse{
		Data: goods,
	})
}

func (handl *Handler) transferMoney(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "could not identify the sender")
		return
	}
	receiverId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id of receiver")
		return
	}
	amount, err := strconv.ParseFloat(ctx.Param("amount"), 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid amount of money")
		return
	}
	err = handl.services.TransferMoney(userId, receiverId, amount)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}