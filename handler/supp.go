package handler

import (
	"Stepuha.net/entities"
	"github.com/gin-gonic/gin"
	"net/http"
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
