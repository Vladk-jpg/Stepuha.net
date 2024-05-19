package handler

import (
	"Stepuha.net/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (handl *Handler) uploadPicture(ctx *gin.Context) {
	goodId, err := strconv.Atoi(ctx.Param("id"))
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = handl.checkBelonging(userId, goodId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	path := fmt.Sprintf("%s/%s", os.Getenv("PICTURE_DIR"), strconv.Itoa(goodId))
	err = os.RemoveAll(path)
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "Error creating directory")
		return
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var filename string
	imgExt := "jpg"
	for header := range form.File {
		filename = header

		arr := strings.Split(filename, ".")
		if len(arr) > 1 {
			imgExt = arr[len(arr)-1]
		}
		continue
	}

	file, _, err := ctx.Request.FormFile(filename)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
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
		return
	}

	fullFileName := fmt.Sprintf("%s.%s", randomFileName(), imgExt)
	fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", path, fullFileName))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
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
		return
	}

	var input entities.UpdateGoodInput
	input.Picture = &fullFileName
	err = handl.services.Update(userId, goodId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
