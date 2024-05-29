package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHandler = "Authorization"
	userCtx              = "userId"
)

func (handl *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHandler)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "Empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	userId, err := handl.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if !handl.checkIfFrozen(userId) {
		newErrorResponse(ctx, http.StatusUnauthorized, "Frozen")
		return
	}

	ctx.Set(userCtx, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

func (handl *Handler) checkBelonging(userId int, goodId int) error {
	err := handl.services.CheckBelonging(userId, goodId)
	if err != nil {
		return err
	}
	return nil
}

func (handl *Handler) checkIfFrozen(userId int) bool {
	isFrozen, err := handl.services.CheckIfFrozen(userId)
	if err != nil {
		return true
	}
	return isFrozen
}
