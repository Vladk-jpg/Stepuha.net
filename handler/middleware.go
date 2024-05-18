package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const (
	authorizationHandler = "Authorization"
	userCtx              = "userId"
	goodCtx              = "goodId"
	pictureCtx           = "picture"
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

func uploadPicture(ctx *gin.Context) string {
	goodId, _ := ctx.Get(goodCtx)
	path := fmt.Sprintf("%s/image/%s", os.Getenv("PICTURE_DIR"), goodId)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error creating directory")
		return ""
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return ""
	}
	var filename string
	imgExt := "jpg"
	for header := range form.File {
		filename = header

		arr := strings.Split(filename, ".")
		if len(arr) > 1 {
			imgExt = arr[len(arr)-1]
		}
	}
	file, _, err := ctx.Request.FormFile(filename)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return ""
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {
			logger.Print("Can't close a file %s", err)
		}
	}(file)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	fullFileName := fmt.Sprintf("%s.%s", randomFileName(), imgExt)
	fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", path, fullFileName))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	defer func(fileOnDisk *os.File) {
		err := fileOnDisk.Close()
		if err != nil {
			logger.Print("can't close a file on disk")
		}
	}(fileOnDisk)

	_, err = fileOnDisk.Write(fileBytes)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	return fmt.Sprintf(fullFileName)
}
