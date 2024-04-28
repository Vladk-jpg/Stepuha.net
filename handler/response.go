package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
)

var buildInfo, _ = debug.ReadBuildInfo()

var loggerParams = zerolog.ConsoleWriter{
	Out:        os.Stderr,
	TimeFormat: time.RFC1123,
	FormatLevel: func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[%s]", i))
	},
	FormatMessage: func(i interface{}) string {
		return fmt.Sprintf("| %s |", i)
	},
	FormatCaller: func(i interface{}) string {
		return filepath.Base(fmt.Sprintf("%s", i))
	},
	PartsExclude: []string{
		zerolog.TimestampFieldName,
	},
}

var logger = zerolog.New(loggerParams).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Caller().
	Int("pid", os.Getpid()).
	Str("go_version", buildInfo.GoVersion).
	Logger()

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(ctx *gin.Context, returnCode int, message string) {
	logger.Print(message)
	ctx.AbortWithStatusJSON(returnCode, errorResponse{message})
}
