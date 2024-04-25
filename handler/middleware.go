package handler

import (
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
	}

	userId, err := handl.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(userCtx, userId)
}
